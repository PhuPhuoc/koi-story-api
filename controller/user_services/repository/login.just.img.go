package userrepository

import (
	"context"
	"fmt"
	"sync"

	usermodel "github.com/PhuPhuoc/koi-story-api/controller/user_services/model"
)

type VerifyResult struct {
	UserID   string
	AvgScore float64
	Verified bool
}

func (store *userStore) LoginWithFaceImage(ctx context.Context, imageURL string) (usermodel.User, error) {
	var user usermodel.User

	query := "select u.id, f.img_url from user u join face_data f on u.id=f.user_id where u.deleted_at is null"
	rows, err := store.db.QueryxContext(ctx, query)
	if err != nil {
		return user, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	users := make(map[string][]string)
	for rows.Next() {
		var id string
		var faceDetectionURL string
		if err := rows.Scan(&id, &faceDetectionURL); err != nil {
			return user, fmt.Errorf("failed to scan row: %w", err)
		}
		users[id] = append(users[id], faceDetectionURL)
	}

	if len(users) == 0 {
		return user, fmt.Errorf("no users found for face detection")
	}

	var highestScoreUser VerifyResult
	highestScoreUser.AvgScore = 0.6 // Set the minimum threshold for login

	// context with timeout and cancel -> stops all goroutines when a user is verified
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	results := make(chan VerifyResult)
	var wg sync.WaitGroup

	// Process each user concurrently
	for userID, faceURLs := range users {
		wg.Add(1)
		go func(id string, faceURLs []string) {
			defer wg.Done()

			var totalScore float64
			var verifiedCount int

			// Verify each of the 3 face images for the user
			for _, faceDetectionURL := range faceURLs {
				verified, threshold, err := store.callFaceDetectionServer(ctx, faceDetectionURL, imageURL)
				if err != nil {
					fmt.Println("Error verifying face:", err)
					return
				}

				if verified {
					totalScore += threshold
					verifiedCount++
				}
			}

			// Calculate the average score of the verified images
			if verifiedCount > 0 {
				avgScore := totalScore / float64(verifiedCount)
				if avgScore > highestScoreUser.AvgScore {
					results <- VerifyResult{
						UserID:   id,
						AvgScore: avgScore,
						Verified: true,
					}
				}
			}
		}(userID, faceURLs)
	}

	// Close the results channel when all goroutines are finished
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results to find the user with the highest score
	for result := range results {
		if result.Verified && result.AvgScore > highestScoreUser.AvgScore {
			highestScoreUser = result
		}
	}

	// Check if a valid user was found with a high enough score
	if highestScoreUser.Verified {
		query := `
			select id, email, display_name, profile_picture_url, user_type
			from user where id = ?
		`
		if err := store.db.Get(&user, query, highestScoreUser.UserID); err != nil {
			return user, fmt.Errorf("failed to query user data: %w", err)
		}
		fmt.Println("userlogin info: ", user)
		fmt.Println("userlogin avg_score: ", highestScoreUser.AvgScore)
		return user, nil
	}

	// If no matching face was detected
	fmt.Println("userlogin info: ", user)
	return user, fmt.Errorf("no matching face detected ~ highscore found: %v", highestScoreUser.AvgScore)
}
