package userrepository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

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

func (store *userStore) callFaceDetectionServer(ctx context.Context, img1, img2 string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 12*time.Second)
	defer cancel()

	reqBody := DFRequest{
		Image1: img1,
		Image2: img2,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return false, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:5005/verify", bytes.NewBuffer(jsonData))
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("deepface service returned non-OK status: %d, body: %s",
			resp.StatusCode, string(respBody))
	}

	var verifyResp map[string]interface{}
	if err := json.Unmarshal(respBody, &verifyResp); err != nil {
		return false, fmt.Errorf("failed to parse response: %w", err)
	}

	// Check if the "verified" field exists and is a boolean
	verified, ok := verifyResp["verified"].(bool)
	if !ok {
		return false, fmt.Errorf("field 'verified' not found")
	}
	return verified, nil
}
