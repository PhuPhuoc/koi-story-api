package model

type UserLoginInfo struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}
