package password

import (
	crypto "golang.org/x/crypto/bcrypt"
)

var (
	defaultBcryptParams = bcryptParams{
		cost: crypto.DefaultCost,
	}
)

type bcryptParams struct {
	cost int
}

func hashWithBcrypt(raw string, params ...bcryptParams) (password, error) {
	p := defaultBcryptParams
	if len(params) > 0 {
		p = params[0]
	}

	hash, err := crypto.GenerateFromPassword([]byte(raw), p.cost)
	if err != nil {
		return nilPassword, err
	}

	pswd := password{
		Algo: argon2id,
		Value: map[string]interface{}{
			"p": p,
			"h": hash,
		},
	}

	return pswd, nil
}
