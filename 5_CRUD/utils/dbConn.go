package utils

import (
	"database/sql"
	"log"

	"github.com/mattn/go-sqlite3"
)

var Db *sql.DB

const (
	createTable = `
	CREATE TABLE IF NOT EXISTS user_info (
    	name  varchar(64) not null check(length(name) > 5),
    	email varchar(256) not null unique primary key check(length(email) > 5) check(email LIKE '%@%'),
    	age integer not null check(age >= 18) 
	);

	CREATE TABLE IF NOT EXISTS bank_account (
    	bank_uuid varchar(256) not null unique,
    	money real default(0),
    	user_email varchar(256) unique not null, 
    	foreign key (user_email) references user_info(email) 
	);`

	DBName   = "goTest"
	PathToDB = "./db/db.db"
)

func MakeConnToDB() {
	sql.Register("sqlite3_with_extensions",
		&sqlite3.SQLiteDriver{
			Extensions: []string{
				"sqlite3_mod_regexp",
			},
		})

	db, err := sql.Open("sqlite3", PathToDB)

	if err != nil {
		log.Fatalln(err)
	}

	Db = db
}

func CreateTables(db *sql.DB) {
	if _, err := db.Exec(createTable); err != nil {
		log.Fatalln(err)
	}

}
