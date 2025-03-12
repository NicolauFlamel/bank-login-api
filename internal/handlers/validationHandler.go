package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/nicolau_flamel/bank-login-api/internal/dtos"
	"github.com/nicolau_flamel/bank-login-api/internal/services"
)

func ValidationHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request){
  return func(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read the body", http.StatusInternalServerError)
        return
    }

    var clientReq dtos.ClientReq

    if err := json.Unmarshal(body, &clientReq); err != nil {
        http.Error(w, "Failed to unmarshal body", http.StatusInternalServerError)
        return
    }

    session, err := services.SessionIsValid(db, clientReq.SessionId)
    if err != nil {
        http.Error(w, fmt.Sprintf("Session is not valid: %v", err), http.StatusInternalServerError)
        return
    }

    err = services.IsSessionExpired(session.CreatedAt)
    if err != nil {
        http.Error(w, fmt.Sprintf("Session is not valid: %v", err), http.StatusInternalServerError)
        return      
    }

    clientReq.Sequence, err = services.DecryptLayoutAESGCM(clientReq.Sequence)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error decrypting sequence: %v", err), http.StatusInternalServerError)
        return        
    }
    
    sequenceLayout, err := services.IsLayoutValid(session.Layout, clientReq.Sequence)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error validating layout: %v", err), http.StatusInternalServerError)
        return        
    }

    fmt.Println(sequenceLayout)

    users := services.GetUsers(db)

    passIsValid, user := services.CheckPasswords(users, sequenceLayout)
    if !passIsValid {
        http.Error(w, fmt.Sprintf("Invalid password"), http.StatusInternalServerError)
        return        
    }

    response := map[string]interface{}{
			"isValid":      passIsValid,
			"user":     user.Name,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		}

  }
}
