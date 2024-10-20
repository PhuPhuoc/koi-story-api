package commenthandler

import (
	"net/http"

	commmentrepository "github.com/PhuPhuoc/koi-story-api/controller/comment_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// GetAllCommentInPost godoc
//
//	@Summary		Get all comment's post
//	@Description	Get all comment's post
//	@Tags			comments
//	@Accept			json
//	@Produce		json
//	@Param			post_id	path		string					true	"Post ID"
//	@Success		200		{object}	map[string]interface{}	"data object"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/post/{post_id}/comment [get]
func getAllCommentHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		post_id := c.Param("post_id")
		if post_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
			return
		}
		repo := commmentrepository.NewCommentStore(db)
		data, err := repo.GetAllCommentInPost(post_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
