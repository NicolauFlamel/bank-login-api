package services

import (
	"database/sql"
	"log"

	"github.com/nicolau_flamel/bank-login-api/internal/dtos"
)

func GetUsers(db *sql.DB) []dtos.User {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []dtos.User
	for rows.Next() {
		var user dtos.User
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
