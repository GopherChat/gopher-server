package store

import "fmt"

///// ErrRecordNotFound /////

type ErrRecordNotFound struct {
	Value string
}

// IsErrRecordNotFound checks if an error is an IsErrRecordNotFound.
func IsErrRecordNotFound(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(ErrRecordNotFound)
	return ok
}

func (err ErrRecordNotFound) Error() string {
	return fmt.Sprintf("record does not exist [value: %s]", err.Value)
}
