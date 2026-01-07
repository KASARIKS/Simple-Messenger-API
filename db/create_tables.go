package db

import "database/sql"

const usersTable = "CREATE TABLE IF NOT EXISTS users (" +
	"id TEXT PRIMARY KEY," +
	"password TEXT NOT NULL," +
	"nickname TEXT NOT NULL" +
	");"

func createTables(db *sql.DB) error {
	if _, err := db.Exec(usersTable); err != nil {
		return err
	}

	return nil
}
