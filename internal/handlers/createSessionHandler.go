package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/nicolau_flamel/bank-login-api/internal/services"
)

func CreateSessionHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request){
  return func(w http.ResponseWriter, r *http.Request) {
    sessionId, err :=  services.GenerateSessionID()
    if err != nil {
      http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		  return
    }
    
    layout, err := services.BuildLayout(db, sessionId)
    if err != nil {
      http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		  return
    }

    encryptedLayout, err := services.EncryptLayoutAESGCM(layout)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error encrypting layout: %v", err), http.StatusInternalServerError)
			return
		}

    response := map[string]interface{}{
			"session_id":      sessionId,
			"layout":     encryptedLayout,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		}
  }
}
