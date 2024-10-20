package markethandler

import (
	"fmt"
	"net/http"

	marketrepository "github.com/PhuPhuoc/koi-story-api/controller/market_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// DeleteMarketPost godoc
//
//	@Summary		Delete market's post
//	@Description	Delete post details by post ID
//	@Tags			markets
//	@Accept			json
//	@Produce		json
//	@Param			post_id	path		string					true	"Post ID"
//	@Success		200		{object}	string					"Deleted successfully"
//	@Failure		404		{object}	map[string]interface{}	"Post not found"
//	@Failure		400		{object}	map[string]interface{}	"Invalid post ID"
//	@Router			/markets/{post_id} [delete]
func deletePostDetailsHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("post_id")
		if postID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
			return
		}

		repo := marketrepository.NewMarketStore(db)

		err := repo.DeleteMarketPost(postID)
		if err != nil {
			fmt.Println("err: ", err)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
	}
}
