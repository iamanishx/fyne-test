package crypto

import (
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/argon2"
)

const (
	Memory      = 64 * 1024
	Iterations  = 3
	Parallelism = 2
	SaltLength  = 16
	KeyLength   = 32
)

func DeriveKey(password, salt []byte) []byte {
	return argon2.IDKey(password, salt, Iterations, Memory, Parallelism, KeyLength)
}

func GenerateSalt() ([]byte, error) {
	b := make([]byte, SaltLength)
	_, err := rand.Read(b)
	if err != nil {
		return nil, errors.New("failed to generate salt")
	}
	return b, nil
}
