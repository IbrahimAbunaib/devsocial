package database

import (
	"database/sql"
	"log"
)

var db *sql.DB

func connnect() {
	var err error
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=MR.ibrahim2001 dbname=devsocial sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to the Database", err)
	}

	// we use ping to verify that the connection is still alive, once disconnected to return error
	err = db.Ping()
	if err != nil {
		log.Fatal("Database unreachable", err)
	}

	log.Println("Succesfully connnected to Database")
}
