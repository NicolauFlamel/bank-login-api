package services

import (
	"encoding/base64"
	"golang.org/x/crypto/argon2"
)

func HashDigit(digit string, salt []byte) string {
	hash := argon2.IDKey([]byte(digit), salt, 1, 64*1024, 4, 16)
	return base64.StdEncoding.EncodeToString(hash)
}

