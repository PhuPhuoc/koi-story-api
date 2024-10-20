package markethandler

import (
	"net/http"

	marketmodel "github.com/PhuPhuoc/koi-story-api/controller/market_services/model"
	marketrepository "github.com/PhuPhuoc/koi-story-api/controller/market_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// CreateNewPostMarket godoc
//
//	@Summary		Update market's post
//	@Description	Update market's post
//	@Tags			markets
//	@Accept			json
//	@Produce		json
//	@Param			post_id	path		string							true	"Post ID"
//	@Param			user	body		marketmodel.UpdatePostMarket	true	"details of current post of market"
//	@Success		200		{object}	map[string]interface{}			"message success"
//	@Failure		400		{object}	error							"Bad request error"
//	@Router			/markets/{post_id} [put]
func updatePostDetailHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		postID := c.Param("post_id")
		if postID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
			return
		}

		var req marketmodel.UpdatePostMarket
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		repo := marketrepository.NewMarketStore(db)

		err := repo.UpdateMarketDetails(postID, req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "updated post successfully"})
	}
}
