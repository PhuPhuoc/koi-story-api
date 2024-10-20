package commenthandler

import (
	"net/http"

	commmentrepository "github.com/PhuPhuoc/koi-story-api/controller/comment_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// DeleteComment godoc
//
//	@Summary		Delete  post'comment
//	@Description	Delete comment in post by image ID
//	@Tags			comments
//	@Accept			json
//	@Produce		json
//	@Param			post_id		path		string					true	"post ID"
//	@Param			comment_id	path		string					true	"comment ID"
//	@Success		204			{object}	string					"Deleted successfully"
//	@Failure		404			{object}	map[string]interface{}	"Post not found"
//	@Failure		400			{object}	map[string]interface{}	"Invalid post ID"
//	@Router			/post/{post_id}/comment/{comment_id} [delete]
func deleteCommentHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		post_id := c.Param("post_id")
		if post_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
			return
		}
		comment_id := c.Param("comment_id")
		if comment_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image ID is required"})
			return
		}

		repo := commmentrepository.NewCommentStore(db)

		err := repo.DeleteComment(comment_id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{"message": "Deleted successfully"})
	}
}
