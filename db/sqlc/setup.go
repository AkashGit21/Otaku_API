package db

import (
	"database/sql"
	"log"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/otaku_db?sslmode=disable"
)

func MakeConnection(q *Queries) (*Queries, error) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Cannot connect to Database: %v", err)
		return nil, err
	}

	q = New(conn)
	log.Printf("Database Connected successfully")

	return q, nil
}
