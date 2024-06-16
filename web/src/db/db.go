package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	// create connection string
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=cool_app"

	// open database connection
}
