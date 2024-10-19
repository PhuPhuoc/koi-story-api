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
//	@Summary		Create new market's post
//	@Description	Create new market's post
//	@Tags			markets
//	@Accept			json
//	@Produce		json
//	@Param			user	body		marketmodel.NewPostMarket	true	"details of new post of market"
//	@Success		201		{object}	map[string]interface{}		"message success"
//	@Failure		400		{object}	error						"Bad request error"
//	@Router			/markets [post]
func createPostHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req marketmodel.NewPostMarket

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		repo := marketrepository.NewMarketStore(db)

		err := repo.CreateNewMarketPost(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "created new post successfully"})
	}
}
