package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Aragon-Joaquin/Learning-GO/handlers"
	"github.com/Aragon-Joaquin/Learning-GO/utils"
)

var db *sql.DB

func main() {
	db = utils.MakeConnToDB()
	utils.CreateTables(db)

	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln(err)
	}

}
