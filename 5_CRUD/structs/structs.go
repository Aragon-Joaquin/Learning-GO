package db_structs

type User_info struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type Bank_account struct {
	Bank_uuid  string  `json:"bank_uuid"`
	Money      float64 `json:"money"`
	User_email string  `json:"user_email"`
}

type Bank_User struct {
	User_info
	Bank_account
}
