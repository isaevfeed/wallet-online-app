package main

import (
	"log"
	"os"
	"wallet/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error to loading env: %s", err)
	}

	srv := server.NewServer(os.Getenv("SERVER_ADDR"))
	srv.Start()
}
