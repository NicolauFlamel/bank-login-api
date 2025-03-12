package main

import (
  "log"
  "github.com/joho/godotenv"
  "os"
)

func main() {
  if err := godotenv.Load(".env"); err != nil {
    log.Fatal("Error loading .env file")
  }

	server := NewApiServer(":8080", os.Getenv("DBCONN"))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
