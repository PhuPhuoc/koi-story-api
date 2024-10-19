package usermodel

type Register struct {
	Email           string `json:"email"`
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm-password"`
}
