package imagehandler

import (
	"net/http"

	imagerepository "github.com/PhuPhuoc/koi-story-api/controller/image_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// DeleteImagePost godoc
//
//	@Summary		Delete image  post
//	@Description	Delete image in post by image ID
//	@Tags			images
//	@Accept			json
//	@Produce		json
//	@Param			post_id		path		string					true	"post ID"
//	@Param			image_id	path		string					true	"image ID"
//	@Success		204			{object}	string					"Deleted successfully"
//	@Failure		404			{object}	map[string]interface{}	"Post not found"
//	@Failure		400			{object}	map[string]interface{}	"Invalid post ID"
//	@Router			/post/{post_id}/image/{image_id} [delete]
func deleteImageInPostMarketHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		post_id := c.Param("post_id")
		if post_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
			return
		}
		image_id := c.Param("image_id")
		if image_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image ID is required"})
			return
		}

		repo := imagerepository.NewImageStore(db)

		err := repo.DeleteImagePost(image_id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
	}
}
