package filter

import (
	"crypto/sha256"
)

type HashFunction func(value string) ([]byte, error)

func Sha256(value string) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write([]byte(value))
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
