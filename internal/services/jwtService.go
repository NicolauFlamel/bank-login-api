package services

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nicolau_flamel/bank-login-api/internal/models"
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

func GenerateJWT(sessionID string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &models.CustomClaims{
		SessionID: sessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign the token: %v", err)
	}

	return signedToken, nil
}

func ValidateJWT(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
		}
		return jwtSecret, nil
	})

	if err != nil {
		return false, fmt.Errorf("invalid token: %v", err)
	}

	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		if claims.ExpiresAt.Before(time.Now()) {
			return false, fmt.Errorf("token has expired")
		}
		return true, nil
	}

	return false, fmt.Errorf("invalid claims")
}
