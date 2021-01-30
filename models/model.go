package models

type User struct {
	FirstName string `json:"first_name,validate:required"`
	LastName  string `json:"last_name,validate:required"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password,validate:required"`
}
