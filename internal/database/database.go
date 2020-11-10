package database

import (
	"database/sql"
	"fmt"
)

// DBConn database connection
func DBConn(host, port, username, password, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, fmt.Errorf("could not open postgresql connection: %w", err)
	}

	return db, err
}
