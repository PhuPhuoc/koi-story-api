package markethandler

import (
	"fmt"
	"net/http"

	marketrepository "github.com/PhuPhuoc/koi-story-api/controller/market_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// GetMarketPostDetails godoc
//
//	@Summary		Get details of a market's post
//	@Description	Get market post details by post ID
//	@Tags			markets
//	@Accept			json
//	@Produce		json
//	@Param			post_id	path		string						true	"Post ID"
//	@Success		200		{object}	marketmodel.MarketDetails	"market post details"
//	@Failure		404		{object}	map[string]interface{}		"Post not found"
//	@Failure		400		{object}	map[string]interface{}		"Invalid post ID"
//	@Router			/markets/{post_id} [get]
func getPostDetailsHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("post_id")
		if postID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
			return
		}

		repo := marketrepository.NewMarketStore(db)

		marketDetails, err := repo.GetMarketDetailsByID(postID)
		if err != nil {
			fmt.Println("err: ", err)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": marketDetails})
	}
}
