package scripts

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

// EncryptPassword encrypts a password using AES encryption and a key
func Encrypt(password, key string) (string, error) {
	// Create a new AES cipher block using the provided key
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Generate a random IV (initialization vector)
	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		return "", err
	}

	// Pad the password to be a multiple of the block size
	paddedPassword := pad([]byte(password), aes.BlockSize)

	// Encrypt the password
	ciphertext := make([]byte, len(paddedPassword))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedPassword)

	// Prepend the IV to the ciphertext (since the IV is needed for decryption)
	encrypted := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// DecryptPassword decrypts an encrypted password using AES and the key
func Decrypt(encryptedPassword, key string) (string, error) {
	// Decode the base64 encrypted password
	data, err := base64.StdEncoding.DecodeString(encryptedPassword)
	if err != nil {
		return "", err
	}

	// Create a new AES cipher block using the provided key
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Extract the IV (first 16 bytes) and ciphertext
	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	// Decrypt the password
	mode := cipher.NewCBCDecrypter(block, iv)
	decoded := make([]byte, len(ciphertext))
	mode.CryptBlocks(decoded, ciphertext)

	// Unpad the decoded password and return it as a string
	unpaddedPassword := unpad(decoded)
	return string(unpaddedPassword), nil
}

func pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func unpad(data []byte) []byte {
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}

