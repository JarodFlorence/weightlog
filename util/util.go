package util

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
	"log"
	"os"
)

var db_url string

func init() {
	db_url = os.Getenv("DATABASE_URL")
	if db_url == "" {
		log.Fatalf("Database connection string is missing")
	}
}

type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func GetDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", db_url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func GetLastId(db DB, tablename string) (id int64, err error) {
	row := db.QueryRow("SELECT currval(pg_get_serial_sequence($1, 'id'));", tablename)
	err = row.Scan(&id)
	return
}
