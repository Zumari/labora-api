package config

import (
	"fmt"
	"net/http"
	"time"
)

func StartServer(router http.Handler) error {
	port := ":9000"
	server := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("server initialized on port: %s... \n", port)
	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("Error while starting up server: '%v'", err)
	}
	return nil

}
