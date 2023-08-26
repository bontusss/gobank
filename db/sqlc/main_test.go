package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/bontusss/gobank/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	testDB, err = sql.Open(config.DBDRIVER, config.DBSOURCE)
	if err != nil {
		log.Fatal("Cannot connect database", err)
	}

	testQueries = New(testDB)

	// start running tests
	os.Exit(m.Run())
}
