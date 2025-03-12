package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"github.com/nicolau_flamel/bank-login-api/internal/models"
)

var key []byte = []byte("myverysecurekeywhichis32byteslon")

func EncryptLayoutAESGCM(plaintextLayout models.Layout) (string, error) {
	plaintext, err := json.Marshal(plaintextLayout)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)

	result := append(nonce, ciphertext...)
	
	return base64.StdEncoding.EncodeToString(result), nil
}

func DecryptLayoutAESGCM(encrypted string) (models.Layout, error) {
	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return models.Layout{}, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return models.Layout{}, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return models.Layout{}, err
	}

	nonceSize := aesGCM.NonceSize()

	if len(data) < nonceSize {
		return models.Layout{}, fmt.Errorf("invalid ciphertext size")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return models.Layout{}, err
	}

	var decryptedLayout models.Layout
	if err := json.Unmarshal(plaintext, &decryptedLayout); err != nil {
		return models.Layout{}, err
	}

	return decryptedLayout, nil
}
