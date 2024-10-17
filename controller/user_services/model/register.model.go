package model

type Register struct {
	Email           string `db:"email" json:"email"`
	Password        string `db:"password" json:"password"`
	ConfirmPassword string `json:"confirm-password"`
}
