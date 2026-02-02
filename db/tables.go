package db

import (
	"database/sql"
	"fmt"
)

func CreateTables(db *sql.DB) error {
	createUsrTable := `
	CREATE TABLE IF NOT EXISTS users(
		id VARCHAR(100) PRIMARY KEY,
		username VARCHAR(100),
		password VARCHAR(100),
		phone VARCHAR(100),
		email VARCHAR(100),
		role VARCHAR(100),
		status VARCHAR(100),
		created_at DATE
	)
	`

	_, err := db.Exec(createUsrTable)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	return nil
}
