package usermodel

type Register struct {
	Email           string `json:"email"`
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	FaceImg         string `json:"face_image"`
}

type RegisterV2 struct {
	Email           string   `json:"email"`
	UserName        string   `json:"user_name"`
	Password        string   `json:"password"`
	ConfirmPassword string   `json:"confirm_password"`
	FaceImg         []string `json:"face_image"`
}
