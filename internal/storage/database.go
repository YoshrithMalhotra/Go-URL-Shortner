package storage

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Connect() {

	connStr := "postgres://postgres:postgres@localhost:5433/urlshortener?sslmode=disable"
	var err error

	DB, err = sql.Open("pgx", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to PostgreSQL ")
}
