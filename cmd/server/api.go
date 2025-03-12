package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/nicolau_flamel/bank-login-api/internal/handlers"
)

const dbKey = "db"

type ApiServer struct {
	addr   string
	dbConn string
}

func NewApiServer(addr string, db string) *ApiServer {
	return &ApiServer{
		addr:   addr,
		dbConn: db,
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permite qualquer origem
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Responder imediatamente a requisições OPTIONS
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}


func (s *ApiServer) Run() error {
  db, err := sql.Open("postgres",os.Getenv("DBCONN"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

  SetupDbUsers(db, true)  

  SetupDbLayouts(db)

	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()

  router.HandleFunc("GET /create-session", handlers.CreateSessionHandler(db))
	router.HandleFunc("POST /validate", handlers.ValidationHandler(db))

	server := http.Server{
		Addr:    s.addr,
		Handler:  corsMiddleware(router),
	}

	fmt.Printf("Server has started on port %s", s.addr)

	return server.ListenAndServe()
}























