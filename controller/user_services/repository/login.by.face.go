package userrepository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// DeepFaceResponse struct to parse the API response
type DeepFaceResponse struct {
	Verified bool    `json:"verified"`
	Distance float64 `json:"distance"`
	Message  string  `json:"message,omitempty"`
}

// DeepFaceRequest struct for the API request
type DeepFaceRequest struct {
	Image1 string `json:"img1_path"`
	Image2 string `json:"img2_path"`
}

// LoginFace performs face verification for user login
func (store *userStore) LoginFace(ctx context.Context, email, base64LoginImage string) error {
	if email == "" || base64LoginImage == "" {
		return fmt.Errorf("email and login image are required")
	}

	// Get stored image from database with timeout context
	var storedbase64 string
	query := `
		SELECT face_detection_data
		FROM user
		WHERE email = ? AND deleted_at IS NULL
	`

	if err := store.db.GetContext(ctx, &storedbase64, query, email); err != nil {
		return fmt.Errorf("failed to get stored image: %w", err)
	}

	if storedbase64 == "" {
		return fmt.Errorf("no stored image found for user")
	}

	// Call deepface service for verification
	verifyResp, err := store.callDeepFaceService(ctx, storedbase64, base64LoginImage)
	if err != nil {
		return fmt.Errorf("face verification failed: %w", err)
	}

	if !verifyResp.Verified {
		return fmt.Errorf("face verification failed: faces don't match (distance: %f)", verifyResp.Distance)
	}

	return nil
}

func (store *userStore) callDeepFaceService(ctx context.Context, storedImage, loginImage string) (*DeepFaceResponse, error) {
	// Create timeout context
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	// Prepare request payload
	reqBody := DeepFaceRequest{
		Image1: "image/jpeg;base64,/" + storedImage,
		Image2: "image/jpeg;base64,/" + loginImage,
	}

	fmt.Println("reqBody: ", reqBody)

	// Marshal request to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:5005/verify", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set correct content type for JSON
	req.Header.Set("Content-Type", "application/json")

	// Send request using http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read and parse response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("deepface service returned non-OK status: %d, body: %s",
			resp.StatusCode, string(respBody))
	}

	var verifyResp DeepFaceResponse
	if err := json.Unmarshal(respBody, &verifyResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &verifyResp, nil
}
