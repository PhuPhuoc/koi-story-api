package userrepository

import (
	"fmt"
)

func (store *userStore) SaveProfileImage(user_id, base64Image string) error {
	query := "update user set face_detection_data = ? where id = ?"
	_, err := store.db.Exec(query, base64Image, user_id)
	if err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}
	return nil
}
