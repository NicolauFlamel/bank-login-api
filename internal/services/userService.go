package services

import (
	"database/sql"
	"log"

	"github.com/nicolau_flamel/bank-login-api/internal/models"
)

func GetUsers(db *sql.DB, user string) []models.User {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Digit1, &user.Digit2, &user.Digit3, &user.Digit4, &user.Salt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

  return users
}
