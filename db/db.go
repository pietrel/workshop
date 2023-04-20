package db

import (
	"github.com/jmoiron/sqlx"
	"log"
	"second/config"
)

var schema = `
CREATE TABLE users (
    id serial primary key,
    first_name text,
    last_name text,
    email text
)`

func InitPg() *sqlx.DB {
	configuration := config.GetDbConfig()
	db, err := sqlx.Connect("postgres", configuration)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func InitDB() {
	configuration := config.GetDbConfig()
	db, err := sqlx.Connect("postgres", configuration)
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(schema)
}
