package services

import (
	"database/sql"
	"fmt"
	"log"
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
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no layout found for session id %s", sessionId)
		}
		return "", err
	}
	return layout, nil
}
