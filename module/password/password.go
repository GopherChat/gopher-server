package password

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/vmihailenco/msgpack/v5"
)

// Passworder ...
type Passworder interface {
	Generate(raw string) (string, error)
	Compare(raw, encHash string) (bool, error)
}

// Algo ...
type Algo uint

const (
	bcrypt Algo = iota
	argon2id
)

func (p Algo) String() string {
	return [...]string{"bcrypt", "argon2id"}[p]
}

type password struct {
	Algo  Algo                   `msgpack:"a,omitempty"`
	Value map[string]interface{} `msgpack:"v,omitempty"`
}

var nilPassword password

func (p password) String() string {
	buff := new(bytes.Buffer)
	wc := base64.NewEncoder(base64.RawStdEncoding, buff)
	_ = msgpack.NewEncoder(wc).SetSortMapKeys(true).Encode(p)
	_ = wc.Close()
	return buff.String()
}

// Generator ...
type Generator struct {
	algo Algo
}

// NewGenerator ...
func NewGenerator(algo Algo) Generator {
	return Generator{algo: algo}
}

// Generate ...
func (pg Generator) Generate(raw string) (string, error) {
	var (
		p   password
		err error
	)

	switch pg.algo {
	case argon2id:
		p, err = hashWithArgon2id(raw)
	default:
		err = fmt.Errorf("invalid password algorithm")
	}

	if err != nil {
		return "", err
	}

	return p.String(), nil
}

// Compare ...
func (pg Generator) Compare(raw, encHash string) {
	panic("not implemented")
}
