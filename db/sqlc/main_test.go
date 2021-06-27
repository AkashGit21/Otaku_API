package db

import (
	"database/sql"
	"log"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Cannot connect to Database: %v", err)
		return
	}
	testQueries = New(conn)

	m.Run()
}
