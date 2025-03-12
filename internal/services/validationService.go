package services

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/nicolau_flamel/bank-login-api/internal/dtos"
)

func SessionIsValid(db *sql.DB, sessionId string) (dtos.Session, error) {
	session, err := GetValidSession(db,sessionId)
	if err != nil {
		return dtos.Session{}, err
	}
	return session, nil
}

func IsSessionExpired(sessionExpiration time.Time) error {
	now := time.Now()

	if sessionExpiration.Before(now.Add(-5 * time.Minute)) {
		return errors.New("Session expired.")
	}
	return nil
}

func IsLayoutValid(session string, sequence string) ([][]int, error) {
	sessionLayout, err := ParseLayout(session)
	if err != nil {
		return  nil, err
	}

	sequenceLayout, err := ParseLayout(sequence)
	if err != nil {
		return  nil, err
	}
	
	isValid := IsSequenceInSession(sessionLayout, sequenceLayout)

	if isValid {
		return sequenceLayout, nil
	}

	return nil, errors.New("Sequence does not match sessions's layout")
}

func VerifyPassword(user dtos.User, layout [][]int) bool {
	if len(layout) != 4  {
		return false 
	}

	storedHashes := []string{user.Digit1, user.Digit2, user.Digit3, user.Digit4}

	for i := 0; i < 4; i++ {
		num1 := HashDigit(fmt.Sprintf("%d",layout[i][0]), []byte(user.Salt))
		num2 := HashDigit(fmt.Sprintf("%d",layout[i][1]), []byte(user.Salt))

		if num1 != storedHashes[i] && num2 != storedHashes[i] {
			return false
		}
	}

	return true
}

func CheckPasswords(users []dtos.User, layout [][]int) (bool, *dtos.User) {
	for _, user := range users {
		if VerifyPassword(user, layout) {
			return true, &user
		}
	}
	return false, nil
}