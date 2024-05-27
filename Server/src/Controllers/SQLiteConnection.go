package controllers

//TODO create correct close of DB method

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitSQLiteCon() {
	// Connect to SQLite
	var err error
	db, err = sql.Open("sqlite3", "./rps.db")
	if err != nil {
		log.Fatal(err)
	}

	prepareDB()
	if err != nil {
		log.Fatal(err)
	}
}

func prepareDB() error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT UNIQUE, password TEXT, funds REAL)")
	if err != nil {
		return err
	}
	statement.Exec()

	statement, err = db.Prepare("CREATE TABLE IF NOT EXISTS transactions (id INTEGER PRIMARY KEY, sender_user_id INTEGER, recipient_user_id INTEGER, amount REAL, FOREIGN KEY(sender_user_id) REFERENCES users(id), FOREIGN KEY(recipient_user_id) REFERENCES users(id))")
	if err != nil {
		return err
	}
	statement.Exec()

	statement, err = db.Prepare("CREATE TABLE IF NOT EXISTS pending_challenges (id INTEGER PRIMARY KEY, sender_user_id INTEGER, recipient_user_id INTEGER, bet REAL, move INTEGER, FOREIGN KEY(sender_user_id) REFERENCES users(id), FOREIGN KEY(recipient_user_id) REFERENCES users(id))")
	if err != nil {
		return err
	}
	statement.Exec()

	return nil
}
