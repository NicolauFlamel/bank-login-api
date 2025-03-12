package services

import (
	"crypto/rand"
	"encoding/hex"
)

var jwtSecret []byte = []byte("sfdlfjalsdf")

func GenerateSessionID() (string, error) {
	numBytes := 40

	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	sessionID := hex.EncodeToString(randomBytes)

	return sessionID, nil
}