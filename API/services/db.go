package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DbConnection contiene un puntero a la base de datos SQL.
type DbConnection struct {
	*sql.DB
}

var Db DbConnection

// const (
// 	host        = "localhost"
// 	port        = "5432"
// 	dbName      = "labora-proyect-1"
// 	rolName     = "postgres1"
// 	rolPassword = "postgres1"
// )

func GetEnvCredentials() (string, string, string, string, string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("host")
	port := os.Getenv("port")
	dbName := os.Getenv("dbName")
	rolName := os.Getenv("rolName")
	rolPassword := os.Getenv("rolPassword")

	return host, port, dbName, rolName, rolPassword

}

func (db *DbConnection) PingOrDie() {
	if err := db.Ping(); err != nil {
		log.Fatalf("Cannot reach database, error: %v", err)
	}
}

func Connect_DB() {
	host, port, dbName, rolName, rolPassword := GetEnvCredentials()
	fmt.Printf("El pinchi HOST WQUEDA COMO: %s", host)
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
