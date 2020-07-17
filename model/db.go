package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	db, err := sql.Open("postgres",
		"host=postgres port=5432 user=postgres "+
			"password=password dbname=db sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
