package main

import (
	"httpTest/handlers"
	"httpTest/utils"

	"log"
	"net/http"
)

func main() {
	utils.MakeConnToDB()
	utils.CreateTables(utils.Db)

	defer utils.Db.Close()

	//mux can do pattern matching with regex
	mux := http.NewServeMux()

	mux.HandleFunc("GET /home", handlers.RootHandler)
	mux.HandleFunc("/users", handlers.Users)
	mux.HandleFunc("/bank", handlers.Bank)

	// 404
	http.HandleFunc("/", handlers.NotFound)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln(err)
	}

}
