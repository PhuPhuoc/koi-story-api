package userrepository

import (
	"context"
	"fmt"
	"sync"
	"time"

	usermodel "github.com/PhuPhuoc/koi-story-api/controller/user_services/model"
)

type AuthResult struct {
	FlagCheckAuth bool
	Threshold     float64
}

func (store *userStore) LoginWithEmailAndFace(ctx context.Context, req usermodel.UserLoginImage) (usermodel.User, error) {
	var userlogin usermodel.User
	if req.Email == "" || req.FilePath == "" {
		return userlogin, fmt.Errorf("email and login image are required")
	}

	query := `
		SELECT id, email, display_name, profile_picture_url, user_type
		FROM user
		WHERE email = ? AND deleted_at IS NULL
	`

	if err := store.db.Get(&userlogin, query, req.Email); err != nil {
		return userlogin, fmt.Errorf("user not found: %w", err)
	}

	var faceImages []string
	faceQuery := `SELECT img_url FROM face_data WHERE user_id = ?`
	if err := store.db.Select(&faceImages, faceQuery, userlogin.Id); err != nil {
		return userlogin, fmt.Errorf("error retrieving face data: %w", err)
	}

	if len(faceImages) == 0 {
		return userlogin, fmt.Errorf("this account isn't registered for face detection")
	}

	resultCh := make(chan bool, len(faceImages))
	var wg sync.WaitGroup
	var authSuccess bool

	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	for _, faceImage := range faceImages {
		wg.Add(1)
		go func(imageInDB string) {
			defer wg.Done()
			flag_check_auth, _, err := store.callFaceDetectionServer(ctx, imageInDB, req.FilePath)
			if err != nil {
				fmt.Printf("error when detecting face: %v\n", err)
				resultCh <- flag_check_auth
				return
			}
			// Gửi kết quả vào channel
			resultCh <- flag_check_auth
		}(faceImage)
	}

	// Chờ tất cả goroutine hoàn tất
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for result := range resultCh {
		if result {
			authSuccess = true
			break
		}
	}

	if authSuccess {
		return userlogin, nil
	} else {
		return userlogin, fmt.Errorf("authentication result: incorrect")
	}
}
