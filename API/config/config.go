package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// StartServer initializes the server to use in the port that we indicate.
func StartServer(router http.Handler) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("portServer")
	servidor := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Starting Server on port %s...\n", port)
	if err := servidor.ListenAndServe(); err != nil {
		return fmt.Errorf("Error while starting up Server: '%v'", err)
	}
	return nil
}
