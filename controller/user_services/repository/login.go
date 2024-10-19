package userrepository

import (
	"fmt"

	usermodel "github.com/PhuPhuoc/koi-story-api/controller/user_services/model"
)

func (store *userStore) Login(user_info usermodel.UserLoginInfo) (usermodel.User, error) {
	var user usermodel.User
	query := "select id, email, password, display_name, profile_picture_url, user_type from user where email = ? and deleted_at is null"
	err := store.db.Get(&user, query, user_info.Email)
	if err != nil {
		return user, fmt.Errorf("can not log in: %v", err)
	}

	if user_info.Password != user.Password {
		return user, fmt.Errorf("wrong password")
	}
	return user, nil
}
