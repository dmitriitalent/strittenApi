package entity

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password_hash"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
}
