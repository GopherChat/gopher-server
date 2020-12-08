package password

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/GopherChat/gopher-server/module/generate"
	"golang.org/x/crypto/argon2"
)

var (
	errInvalidHash         = errors.New("the encoded hash is not in the correct format")
	errIncompatibleVersion = errors.New("incompatible version of argon2")
	defaultArgonParams     = &params{
		memory:      64 * 1024,
		iterations:  1,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func hashWithArgon2id(raw string, params ...*params) (password, error) {
	p := defaultArgonParams
	if len(params) > 0 {
		p = params[0]
	}

	salt, err := generate.RandomBytes(p.saltLength)
	if err != nil {
		return nilPassword, err
	}

	hash := argon2.IDKey([]byte(raw), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	pswd := password{
		Algo: argon2id,
		Value: map[string]interface{}{
			"v": argon2.Version,
			"m": p.memory,
			"t": p.iterations,
			"p": p.parallelism,
			"s": salt,
			"h": hash,
		},
	}

	return pswd, nil
}

func ComparePasswordAndArgon2idHash(password, encodedHash string) (bool, error) {
	p, salt, hash, err := decodeArgon2idHash(encodedHash)
	if err != nil {
		return false, err
	}
	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
	return subtle.ConstantTimeCompare(hash, otherHash) == 1, nil
}

func MustComparePasswordAndArgon2idHash(password, encodedHash string) bool {
	match, err := ComparePasswordAndArgon2idHash(password, encodedHash)
	if err != nil {
		panic(err)
	}
	return match
}

func decodeArgon2idHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errIncompatibleVersion
	}

	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}
