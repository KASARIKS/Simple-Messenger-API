package db

import "database/sql"

const usersTable = "CREATE TABLE IF NOT EXISTS users (" +
	"id TEXT PRIMARY KEY," +
	"password TEXT NOT NULL," +
	"nickname TEXT NOT NULL" +
	");"

const messagesTable = "CREATE TABLE IF NOT EXISTS messages (" +
	"id INTEGER PRIMARY KEY AUTOINCREMENT," +
	"authorId TEXT NOT NULL," +
	"recipientId TEXT NOT NULL," +
	"value TEXT NOT NULL," +
	"createdAt TEXT NOT NULL" +
	"FOREIGN KEY (authorId) REFERENCES (users)," +
	"FOREIGN KEY (recipientId) REFERENCES (users)" +
	");"

func createTables(db *sql.DB) error {
	if _, err := db.Exec(usersTable); err != nil {
		return err
	}

	if _, err := db.Exec(messagesTable); err != nil {
		return err
	}

	return nil
}
