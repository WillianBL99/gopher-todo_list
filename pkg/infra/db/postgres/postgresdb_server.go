package postgres

import (
	"database/sql"
	"fmt"
)

func ConectDB() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	defer closeDB(db)

	if err != nil {
		panic(err.Error())
	}

	return db;
}

func closeDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connection closed")
	}
}