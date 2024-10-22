package userhandler

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	userrepository "github.com/PhuPhuoc/koi-story-api/controller/user_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// LoginFace godoc
//
//	@Summary		Login by user face
//	@Description	Login by user face
//	@Tags			users
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			file	formData	file					true	"User profile image"
//	@Param			email	formData	string					true	"User email"
//	@Success		200		{object}	map[string]interface{}	"Image uploaded successfully"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/users/loginbyface [post]
func loginFaceHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image"})
			return
		}

		// Đọc địa chỉ email từ form data
		email := c.PostForm("email")
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
			return
		}

		// Mở file
		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to open file"})
			return
		}
		defer src.Close()

		// Đọc nội dung file hình ảnh
		buffer, err := io.ReadAll(src)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read image"})
			return
		}

		// Chuyển đổi hình ảnh sang định dạng base64
		base64Image := base64.StdEncoding.EncodeToString(buffer)

		repo := userrepository.NewUserStore(db)

		// Gọi phương thức đăng nhập bằng khuôn mặt
		err = repo.LoginFace(c, email, base64Image)
		if err != nil {
			fmt.Println("Err: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to login"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login successfully"})
	}
}
