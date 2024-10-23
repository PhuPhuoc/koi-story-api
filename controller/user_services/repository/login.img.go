package userrepository

import (
	"context"
	"fmt"

	usermodel "github.com/PhuPhuoc/koi-story-api/controller/user_services/model"
)

type DFRequest struct {
	Image1 string `json:"img1_path"`
	Image2 string `json:"img2_path"`
}

func (store *userStore) LoginFaceImage(ctx context.Context, req usermodel.UserLoginImage) (usermodel.User, error) {
	var userlogin usermodel.User
	if req.Email == "" || req.FilePath == "" {
		return userlogin, fmt.Errorf("email and login image are required")
	}

	query := `
		SELECT id, email, display_name, profile_picture_url, user_type, face_detection_data
		FROM user
		WHERE email = ? AND deleted_at IS NULL
	`

	if err := store.db.Get(&userlogin, query, req.Email); err != nil {
		return userlogin, fmt.Errorf("user not found: %w", err)
	}

	if userlogin.FaceDetectionData == "" {
		return userlogin, fmt.Errorf("this account isn't register face detection")
	}

	flag_check_auth, err := store.callFaceDetectionServer(ctx, userlogin.FaceDetectionData, req.FilePath)
	if err != nil {
		return userlogin, fmt.Errorf("error when detecting face: %w", err)
	}

	if flag_check_auth {
		return userlogin, nil
	} else {
		return userlogin, fmt.Errorf("authentication result: incorrect")
	}
}
