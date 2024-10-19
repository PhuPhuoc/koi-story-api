package markethandler

import (
	"net/http"

	marketrepository "github.com/PhuPhuoc/koi-story-api/controller/market_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// GetAllMarketPost godoc
//
//	@Summary		Get all market's post
//	@Description	Get all market's post
//	@Tags			markets
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"data object"
//	@Failure		400	{object}	error					"Bad request error"
//	@Router			/markets [get]
func getAllPostHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo := marketrepository.NewMarketStore(db)
		data, err := repo.GetAllPost()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
