package services

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

var historySize int = 100

func BuildLayout(db *sql.DB, sessionId string) (string, error) {
	var layout string
	for {
	  layout := generateLayout()

		fmt.Println(layout)

		uniqueLayout := IsUniqueLayout(db, layout)

		if uniqueLayout {
			InsertLayout(db, layout, sessionId, true)
			break
		}
	}

	return layout, nil
}

func generateLayout() string {
	numbers := rand.Perm(10)
	var keys [5][2]int
	var result strings.Builder

	for i := 0; i < 5; i++ {
		keys[i][0] = numbers[i*2]
		keys[i][1] = numbers[i*2+1]
		if i > 0 {
			result.WriteString(";")
		}
		result.WriteString(fmt.Sprintf("%d,%d", keys[i][0], keys[i][1]))
	}

	return result.String()
}

func ParseLayout(layoutStr string) ([][]int, error) {
	var layout [][]int
	pairs := strings.Split(layoutStr, ";")

	for i, pair := range pairs {
		numbers := strings.Split(pair, ",")
		if len(numbers) != 2 {
			return nil, fmt.Errorf("invalid pair format at index %d", i)
		}

		n1, err1 := strconv.Atoi(numbers[0])
		n2, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid number format at index %d", i)
		}

		layout = append(layout, []int{n1, n2})
	}

	return layout, nil
}

func IsSequenceInSession(sessionLayout, sequenceLayout [][]int) bool {
	sessionMap := make(map[string]struct{})
	for _, pair := range sessionLayout {
		key := fmt.Sprintf("%d,%d", pair[0], pair[1])
		sessionMap[key] = struct{}{}
	}

	for _, pair := range sequenceLayout {
		key := fmt.Sprintf("%d,%d", pair[0], pair[1])
		if _, exists := sessionMap[key]; !exists {
			return false
		}
	}

	return true
}