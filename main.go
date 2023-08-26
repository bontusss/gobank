package main

import (
	"database/sql"
	"log"

	"github.com/bontusss/gobank/api"
	db "github.com/bontusss/gobank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/go_bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect database", err)
	}
	store := db.NewStore(conn) 
	server := api.NewServer(store)

	err = server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Error starting server, %s", err)
	}
}
