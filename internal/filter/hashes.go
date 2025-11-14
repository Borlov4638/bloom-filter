package filter

import (
	"crypto/sha256"
	"crypto/sha512"
)

type HashFunction func(value string) ([]byte, error)

func Sha256WithSalt(value, salt string) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write([]byte(value + salt))
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func Sha512WithSalt(value, salt string) ([]byte, error) {
	h := sha512.New()
	_, err := h.Write([]byte(value + salt))
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func GetKHashFunctionos(k int) []HashFunction {
	res := make([]HashFunction, k)

	for i := 0; i < k; i++ {
		genFunction := func(value string) ([]byte, error) {
			switch i % 2 {
			case 0:
				return Sha256WithSalt(value, string(i))
			case 1:
				return Sha512WithSalt(value, string(i))
			default:
				return Sha256WithSalt(value, string(i))
			}
		}
		res[i] = genFunction
	}
	return res
}
