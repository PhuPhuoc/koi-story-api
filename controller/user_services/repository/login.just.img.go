package userrepository

import (
	"context"
	"fmt"
	"sync"

	usermodel "github.com/PhuPhuoc/koi-story-api/controller/user_services/model"
)

type VerifyResult struct {
	UserID   string
	Verified bool
}

func (store *userStore) LoginWithFaceImage(ctx context.Context, imageURL string) (usermodel.User, error) {
	var user usermodel.User

	query := "select id, face_detection_data from user where deleted_at IS NULL"
	rows, err := store.db.QueryxContext(ctx, query)
	if err != nil {
		return user, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	users := make(map[string]string)
	for rows.Next() {
		var id string
		var faceDetectionURL string
		if err := rows.Scan(&id, &faceDetectionURL); err != nil {
			return user, fmt.Errorf("failed to scan row: %w", err)
		}
		users[id] = faceDetectionURL
	}

	if len(users) == 0 {
		return user, fmt.Errorf("no users found for face detection")
	}

	// context với timeout và cancel -> dừng tất cả goroutines khi có 1 user được verify
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	results := make(chan VerifyResult) // Channel để nhận kết quả từ các goroutine
	var wg sync.WaitGroup              // WaitGroup để quản lý đồng bộ goroutine

	// Sử dụng goroutines để gọi đến API deepface đồng thời
	for userID, faceURL := range users {
		wg.Add(1)
		go func(id string, faceDetectionURL string) {
			defer wg.Done()

			verified, err := store.callFaceDetectionServer(ctx, faceDetectionURL, imageURL)
			if err != nil {
				// Log hoặc xử lý lỗi tại đây
				fmt.Println("Error verifying face:", err)
				return
			}

			// Nếu xác thực đúng, gửi kết quả về channel và hủy context để dừng các goroutine khác
			if verified {
				results <- VerifyResult{UserID: id, Verified: true}
				cancel() // Dừng các công việc còn lại
			}
		}(userID, faceURL)
	}

	// Chạy goroutine khác để đóng channel khi tất cả goroutines hoàn thành
	go func() {
		wg.Wait()
		close(results)
	}()

	// Đợi kết quả từ channel
	for result := range results {
		if result.Verified {
			// Nếu có kết quả đúng, truy vấn user và trả về
			query := `
				select id, email, display_name, profile_picture_url, user_type
				from user where id = ?
			`
			if err := store.db.Get(&user, query, result.UserID); err != nil {
				return user, fmt.Errorf("failed to query user data: %w", err)
			}
			return user, nil
		}
	}

	// Nếu không tìm thấy kết quả đúng
	return user, fmt.Errorf("no matching face detected")

}
