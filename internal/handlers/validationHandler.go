package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type userRequest struct {
  username string `json:"username"`
}

func userHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request){
  return func(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read the body", http.StatusInternalServerError)
        return
    }

    var username userRequest

    if err := json.Unmarshal(body, &username); err != nil {
        log.Fatal(err)
    }

    fmt.Println(username.username)
    
  }
}
