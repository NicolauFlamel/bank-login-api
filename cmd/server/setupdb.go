package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"

	"github.com/nicolau_flamel/bank-login-api/internal/scripts"
)

func SetupDbUsers(db *sql.DB, addUsers bool) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			digit1 TEXT,
			digit2 TEXT,
			digit3 TEXT,
			digit4 TEXT,
			salt TEXT
		)
	`)
	if err != nil {
		return err
	}

	if !addUsers {
		return nil
	}

	users := []string{"dragon slayer", "druid", "oni", "abyss walker", "sigil knight"}
	passwords := []string{"1234", "8539", "3050", "2347", "1596"}

	for i := 0; i < len(users); i++ {
		user := users[i]
		password := passwords[i]

		salt := make([]byte, 16)
		_, err := rand.Read(salt)
		if err != nil {
			return err
		}

		var hashedDigits []string
		for _, digit := range password {
			hashedDigits = append(hashedDigits, scripts.HashDigit(string(digit), salt))
		}

		digit1 := hashedDigits[0]
		digit2 := hashedDigits[1]
		digit3 := hashedDigits[2]
		digit4 := hashedDigits[3]

		_, err = db.Exec(`
			INSERT INTO users (name, digit1, digit2, digit3, digit4, salt)
			VALUES ($1, $2, $3, $4, $5, $6)
		`, user, digit1, digit2, digit3, digit4, fmt.Sprintf("%x", salt))

		if err != nil {
			return err
		}
	}
	return nil
}

func SetupDbLayouts(db *sql.DB) {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS layouts (
        id SERIAL PRIMARY KEY,
        layout JSONB NOT NULL,
        session_id TEXT,
        is_valid BOOLEAN,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )
`)
	if err != nil {
		log.Fatal(err)
	}
}

