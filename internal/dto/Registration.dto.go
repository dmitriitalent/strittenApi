package dto

type Registration struct {
	Login           string `json:"login"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Email           string `json:"email"`
}
