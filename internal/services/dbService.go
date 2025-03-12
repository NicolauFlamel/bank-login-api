package services

import (
	"database/sql"
	"errors"
	"log"

	"github.com/nicolau_flamel/bank-login-api/internal/dtos"
)

func IsUniqueLayout(db *sql.DB, layout string) bool {
	var count int
	query := `
		SELECT COUNT(*) 
		FROM layouts 
		WHERE layout = $1 
		AND id IN (
			SELECT id 
			FROM layouts 
			ORDER BY created_at DESC 
			LIMIT 100
		)`
	err := db.QueryRow(query, layout).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count == 0
}

func InsertLayout(db *sql.DB, layout string, sessionId string, isValid bool) {
	_, err := db.Exec("INSERT INTO layouts (layout, session_id, is_valid) VALUES ($1, $2, $3)", layout, sessionId, isValid)
	if err != nil {
		log.Fatal(err)
	}
}

func GetLayoutBySessionID(db *sql.DB, sessionId string) (string, error) {
	var layout string 
	query := "SELECT layout FROM layouts WHERE session_id = $1"
	err := db.QueryRow(query, sessionId).Scan(&layout)
	if err != nil {
		return "", err
	}

	return layout, nil
}

func GetValidSession(db *sql.DB, sessionId string) (dtos.Session, error) {
	var session dtos.Session
	query := "SELECT session_id, layout, is_valid, created_at FROM layouts WHERE session_id = $1"
	err := db.QueryRow(query, sessionId).Scan(&session.Id, &session.Layout, &session.IsValid, &session.CreatedAt)
	if err == sql.ErrNoRows {
		return dtos.Session{}, errors.New("Session token is invalid")
	}

	if session.IsValid == true {
		query := "UPDATE layouts SET is_valid = false WHERE session_id = $1;"
		_, err := db.Exec(query ,sessionId)
		if err != nil {
			return dtos.Session{}, errors.New("Could not update session status")
		}
	} else {
		return dtos.Session{}, errors.New("Session has been invalidated")
	}
	return session, err
}
