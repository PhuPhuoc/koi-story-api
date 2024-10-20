package consulthandler

import (
	"fmt"
	"net/http"

	consultrepository "github.com/PhuPhuoc/koi-story-api/controller/consult_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// DeleteConsultPost godoc
//
//	@Summary		Delete consult's post
//	@Description	Delete post details by post ID
//	@Tags			consults
//	@Accept			json
//	@Produce		json
//	@Param			post_id	path		string					true	"Post ID"
//	@Success		200		{object}	string					"Deleted successfully"
//	@Failure		404		{object}	map[string]interface{}	"Post not found"
//	@Failure		400		{object}	map[string]interface{}	"Invalid post ID"
//	@Router			/consults/{post_id} [delete]
func deletePostDetailsHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("post_id")
		if postID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
			return
		}

		repo := consultrepository.NewConsultStore(db)

		err := repo.DeleteConsultPost(postID)
		if err != nil {
			fmt.Println("err: ", err)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
	}
}
