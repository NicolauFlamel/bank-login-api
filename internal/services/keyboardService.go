package services

import (
	"database/sql"
	"encoding/json"
	"math/rand"
  "github.com/nicolau_flamel/bank-login-api/internal/models"
)

var historySize int = 100

func BuildLayout(db *sql.DB, sessionId string) (models.Layout, error) {
	var resLayout models.Layout  
	for {
	  layout := generateLayout()

		layoutJSON, err := json.Marshal(layout)
		if err != nil {
      return models.Layout{}, err
		}

		uniqueLayout := IsUniqueLayout(db, string(layoutJSON))

		if uniqueLayout {
			InsertLayout(db, string(layoutJSON), sessionId, true)

			resLayout = layout
			break
		}
	}

	return resLayout, nil
}

func generateLayout() models.Layout {
	numbers := rand.Perm(10)
	var keys models.Layout

	for i := 0; i < 5; i++ {
		keys[i][0] = numbers[i*2]
		keys[i][1] = numbers[i*2+1]
	}

	return keys
}

