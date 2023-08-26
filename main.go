package main

import (
	"database/sql"
	"log"

	"github.com/bontusss/gobank/api"
	db "github.com/bontusss/gobank/db/sqlc"
	"github.com/bontusss/gobank/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}
	conn, err := sql.Open(config.DBDRIVER, config.DBSOURCE)
	if err != nil {
		log.Fatal("Cannot connect database", err)
	}
	store := db.NewStore(conn) 
	server := api.NewServer(store)

	err = server.Start(config.PORT)
	if err != nil {
		log.Fatalf("Error starting server, %s", err)
	}
}
