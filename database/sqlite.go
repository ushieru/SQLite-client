package database

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

const InMemory = ":memory:"

func GetDatabase(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", "db.db")
	return db, err
}
