package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // _ is used as we don't use this package directly but we do use it through the sql package 
)

var DB *sql.DB

func init() {
	var err error 
	// keep the function name init when initialising a database so that go understands that this function is the entry point to initialise a database and the function would not need to be called anywhere this way
	DB, err = sql.Open("sqlite3", "api.db") // sql is a built-in go package that lets us communicate with our database. Open takes 2 arguments, the database we want to use and the name of the local file containing our database
	if err != nil {
		panic("could not connect to database")
	}

	DB.SetMaxOpenConns(10) // we are not really making connections when we use Open. Now this function is used to control how many max connections we can make simultaneously
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}
}