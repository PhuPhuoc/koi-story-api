package userrepository

import (
	"fmt"

	usermodel "github.com/PhuPhuoc/koi-story-api/controller/user_services/model"
	"github.com/PhuPhuoc/koi-story-api/utils"
	"github.com/google/uuid"
)

func (store *userStore) RegisterNewAccountV2(user_info usermodel.RegisterV2) error {
	tx, err := store.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	if user_info.Password != user_info.ConfirmPassword {
		return fmt.Errorf("passwords do not match")
	}

	flag_duplicate_email := false
	rawsql_checkEmailExist := `select exists(select 1 from user where email = ?)`

	if err := store.db.Get(&flag_duplicate_email, rawsql_checkEmailExist, user_info.Email); err != nil {
		return fmt.Errorf("unable to check for duplicate emails: %v", err)
	}
	if flag_duplicate_email {
		return fmt.Errorf("email already exists")
	}

	user_id := uuid.New().String()
	newUser := usermodel.User{
		Id:                user_id,
		Email:             user_info.Email,
		DisplayName:       user_info.UserName,
		Password:          user_info.Password,
		ProfilePictureUrl: "https://cdn.dribbble.com/users/113499/screenshots/13947091/media/85c35fe30676eaae21f1b6401d9809b4.png",
		UserType:          "user",
		CreatedAt:         utils.CreateDateTimeCurrentFormated(),
	}

	rawsql_createNewAccount := `
	insert into user (id, email, password, display_name, profile_picture_url, user_type, created_at)
	values (:id, :email, :password, :display_name, :profile_picture_url, :user_type, :created_at)
	`
	_, err = tx.NamedExec(rawsql_createNewAccount, newUser)
	if err != nil {
		return fmt.Errorf("failed to insert new user: %w", err)
	}

	query := `
        INSERT INTO face_data (user_id, img_url) VALUES (?, ?)
    `

	// Thực hiện insert cho từng URL
	stmt, err := tx.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	for _, url := range user_info.FaceImg {
		if _, err := stmt.Exec(user_id, url); err != nil {
			return fmt.Errorf("failed to insert data for URL %s: %v", url, err)
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}
