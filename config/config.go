package config

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host        = "localhost"
	port        = "5432"
	dbName      = "labora-proyect-1"
	rolName     = "postgres1"
	rolPassword = "postgres1"
)

func DBConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, rolName, rolPassword, dbName)
	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful connection to the database:", dbConn)
	return dbConn, err
}
