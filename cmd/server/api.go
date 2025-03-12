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

func (s *ApiServer) Run() error {
  db, err := sql.Open("postgres",os.Getenv("DBCONN"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

  SetupDbUsers(db, false)  

  SetupDbLayouts(db)

	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()

  router.HandleFunc("GET /create-session", handlers.CreateSessionHandler(db))
	router.HandleFunc("POST /validate", handlers.ValidationHandler(db))

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	fmt.Printf("Server has started on port %s", s.addr)

	return server.ListenAndServe()
}























