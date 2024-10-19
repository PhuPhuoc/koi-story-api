package usermodel

type User struct {
	Id                string `db:"id" json:"fb_uid"`
	Email             string `db:"email" json:"email"`
	Password          string `db:"password" json:"-"`
	DisplayName       string `db:"display_name" json:"display_name"`
	ProfilePictureUrl string `db:"profile_picture_url" json:"profile_picture_url"`
	UserType          string `db:"user_type" json:"user_type"`
	CreatedAt         string `db:"created_at" json:"created_at,omitempty"`
	DeletedAt         string `db:"deleted_at" json:"deleted_at,omitempty"`
}
