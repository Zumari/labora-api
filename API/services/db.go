package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DbConnection contiene un puntero a la base de datos SQL.
type DbConnection struct {
	*sql.DB
}

var Db DbConnection

const (
	host        = "localhost"
	port        = "5432"
	dbName      = "labora-proyect-1"
	rolName     = "postgres1"
	rolPassword = "postgres1"
)

func (db *DbConnection) PingOrDie() {
	if err := db.Ping(); err != nil {
		log.Fatalf("Cannot reach database, error: %v", err)
	}
}

func Connect_DB() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, rolName, rolPassword, dbName)
	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful connection to the database:", dbConn)
	Db = DbConnection{dbConn}
	Db.PingOrDie()
	if err != nil {
		log.Fatal(err)
	}
}
