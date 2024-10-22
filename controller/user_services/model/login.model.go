package usermodel

type UserLoginInfo struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type LoginFace struct {
	Email string `db:"email" json:"email"`
}

type UserLoginImage struct {
	Email    string `json:"email"`
	FilePath string `db:"file_path" json:"file_path"`
}
