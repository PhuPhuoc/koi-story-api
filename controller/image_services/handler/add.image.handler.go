package imagehandler

import (
	"net/http"

	imagemodel "github.com/PhuPhuoc/koi-story-api/controller/image_services/model"
	imagerepository "github.com/PhuPhuoc/koi-story-api/controller/image_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// AddNewImageInMarketPost godoc
//
//	@Summary		Add new image to post
//	@Description	Add new image to post
//	@Tags			images
//	@Accept			json
//	@Produce		json
//	@Param			post_id	path		string					true	"Post ID"
//	@Param			url		body		imagemodel.UpdateImage	true	"new url"
//	@Success		201		{object}	map[string]interface{}	"message success"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/post/{post_id}/image [post]
func addImageHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		post_id := c.Param("post_id")
		if post_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
			return
		}

		var req imagemodel.UpdateImage
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}


		repo := imagerepository.NewImageStore(db)

		err := repo.AddNewImageInPost(post_id, req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "add new post'image successfully"})
	}
}
