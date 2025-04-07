package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	db_structs "httpTest/structs"
	"httpTest/utils"
	"log"
	"net/http"

	"github.com/google/uuid"
)

/*
{
	"money": 12300,
	"user_email": "joeDoe@email.com"
}
*/

func Bank(w http.ResponseWriter, r *http.Request) {
	var bank_acc db_structs.Bank_account

	switch r.Method {
	case http.MethodGet:
		{
			if err := json.NewDecoder(r.Body).Decode(&bank_acc); err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var bank_details db_structs.Bank_User
			err := utils.Db.QueryRow(`SELECT b.money, u.name FROM bank_account b INNER JOIN user_info u
			ON b.user_email = u.email WHERE $1 = b.user_email LIMIT 1;`, bank_acc.User_email).Scan(&bank_details.Money, &bank_details.Name)

			fmt.Println(err, bank_details)
			if err != sql.ErrNoRows && err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if (db_structs.Bank_User{}) != bank_details {
				json.NewEncoder(w).Encode(bank_details)
			} else {
				json.NewEncoder(w).Encode(make([]int, 0))
			}

		}
	case http.MethodPost:
		{
			if err := json.NewDecoder(r.Body).Decode(&bank_acc); err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var userIsReal db_structs.User_info
			err := utils.Db.QueryRow(`SELECT email FROM user_info WHERE $1 = email;`,
				bank_acc.User_email).Scan(&userIsReal.Email)

			fmt.Println(userIsReal, bank_acc)
			if err != nil {
				log.Fatalln("User does not exists.")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			id := uuid.New()
			var results db_structs.Bank_account
			err = utils.Db.QueryRow(`INSERT INTO 
			bank_account(bank_uuid, user_email) VALUES($1, $2) RETURNING *;`,
				id, &userIsReal.Email).Scan(&results.Bank_uuid, &results.Money, &results.User_email)

			if err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)

			//! i know this leaks the bank_uuid
			json.NewEncoder(w).Encode(results)
		}
	case http.MethodPatch:
		{
			if err := json.NewDecoder(r.Body).Decode(&bank_acc); err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var newResults db_structs.Bank_account
			err := utils.Db.QueryRow(`
			UPDATE bank_account set money = $1 WHERE user_email = $2 RETURNING money, user_email; 
			`, bank_acc.Money, bank_acc.User_email).Scan(&newResults.Money, &newResults.User_email)

			if err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(newResults)
		}
	//! http.MethodDelete, but i get the idea
	default:
		MethodNotAllowed(w)
	}
}
