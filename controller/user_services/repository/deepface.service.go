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

type DFRequest struct {
	Image1 string `json:"img1_path"`
	Image2 string `json:"img2_path"`
}

func (store *userStore) callFaceDetectionServer(ctx context.Context, img1, img2 string) (bool, float64, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	reqBody := DFRequest{
		Image1: img1,
		Image2: img2,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return false, 0, fmt.Errorf("failed to marshal request: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:5005/verify", bytes.NewBuffer(jsonData))

	if err != nil {
		return false, 0, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, 0, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, 0, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return false, 0, fmt.Errorf("deepface service returned non-OK status: %d, body: %s",
			resp.StatusCode, string(respBody))
	}

	var verifyResp map[string]interface{}
	if err := json.Unmarshal(respBody, &verifyResp); err != nil {
		return false, 0, fmt.Errorf("failed to parse response: %w", err)
	}

	// Check if the "verified" field exists and is a boolean
	verified, ok := verifyResp["verified"].(bool)
	if !ok {
		return false, 0, fmt.Errorf("field 'verified' not found")
	}
	threshold, ok := verifyResp["threshold"].(float64)
	if !ok {
		return false, 0, fmt.Errorf("field 'threshold' not found")
	}
	return verified, threshold, nil
}
