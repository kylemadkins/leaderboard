package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	PORT = 5432
)

var ErrNoMatch = fmt.Errorf("No matching record")

type Database struct {
	Conn *sql.DB
}

func Initialize(url string) (Database, error) {
	db := Database{}
	conn, err := sql.Open("postgres", url)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}
