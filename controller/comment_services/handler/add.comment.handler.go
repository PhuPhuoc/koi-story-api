package commenthandler

import (
	"net/http"

	commentmodel "github.com/PhuPhuoc/koi-story-api/controller/comment_services/model"
	commmentrepository "github.com/PhuPhuoc/koi-story-api/controller/comment_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// AddNewCommentInPost godoc
//
//	@Summary		Add new commnet to post
//	@Description	Add new comment to post
//	@Tags			comments
//	@Accept			json
//	@Produce		json
//	@Param			post_id	path		string							true	"Post ID"
//	@Param			comment	body		commentmodel.CreateCommentModel	true	"comment info"
//	@Success		201		{object}	map[string]interface{}			"message success"
//	@Failure		400		{object}	error							"Bad request error"
//	@Router			/post/{post_id}/comment [post]
func addCommentHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		post_id := c.Param("post_id")
		if post_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
			return
		}

		var req commentmodel.CreateCommentModel

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		repo := commmentrepository.NewCommentStore(db)

		err := repo.AddCommentToPost(post_id, req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "comment successfully"})
	}
}
