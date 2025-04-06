package handlers

import (
	"encoding/json"
	db_structs "httpTest/structs"
	"httpTest/utils"
	"log"
	"net/http"
)

/* //! http://localhost:8080/users
Send this JSON using ThunderClient, PostMan or using cURL in the cli
{
	"name": "joe doe",
	"age": 21,
	"email": "joeDoe@email.com"
}

*/

func Users(w http.ResponseWriter, r *http.Request) {
	var user db_structs.User_info

	switch r.Method {
	case http.MethodGet:
		{
			res, err := utils.Db.Query(`SELECT * FROM user_info;`)

			if err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var totalUsers = make([]db_structs.User_info, 0)
			for res.Next() {

				err := res.Scan(&user.Name, &user.Email, &user.Age)

				if err != nil {
					log.Fatalln(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				totalUsers = append(totalUsers, user)
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(totalUsers)
		}

	case http.MethodPost:
		{

			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var user_inserted string
			err := utils.Db.QueryRow(`
				INSERT INTO user_info(name,email,age) 
				VALUES($1, $2, $3) RETURNING name;`,
				&user.Name, &user.Email, &user.Age).Scan(&user_inserted)

			if err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(user_inserted)
		}

	case http.MethodPatch:
		{
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var users_updated db_structs.User_info
			err := utils.Db.QueryRow(`UPDATE user_info SET age = $1, name = $2 WHERE $3 = email RETURNING name,age`, &user.Age, &user.Name, &user.Email).Scan(&users_updated.Name, &users_updated.Age)

			if err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(users_updated)
		}

	case http.MethodDelete:
		{
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var user_deleted string
			err := utils.Db.QueryRow(`DELETE FROM user_info WHERE $1 = email RETURNING email`, &user.Email).Scan(&user_deleted)

			if err != nil {
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(user_deleted)
		}
	default:
		w.Write([]byte("Method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)

	}

}
