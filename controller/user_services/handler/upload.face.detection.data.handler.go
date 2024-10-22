package userhandler

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"

	userrepository "github.com/PhuPhuoc/koi-story-api/controller/user_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// UploadImage godoc
//
//	@Summary		Upload user profile image
//	@Description	Upload and save user profile image as base64
//	@Tags			users
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			user_id	path		string					true	"User ID"
//	@Param			file	formData	file					true	"User profile image"
//	@Success		200		{object}	map[string]interface{}	"Image uploaded successfully"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/users/{user_id}/uploadimage [post]
func uploadImageHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		if user_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}

		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image"})
			return
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read image"})
			return
		}

		base64Image := base64.StdEncoding.EncodeToString(fileBytes)

		repo := userrepository.NewUserStore(db)

		err = repo.SaveProfileImage(user_id, base64Image)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to save image"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
	}
}
