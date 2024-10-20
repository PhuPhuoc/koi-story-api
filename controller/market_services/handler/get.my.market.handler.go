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
//	@Summary		Get all my market's post
//	@Description	Get all my market's post
//	@Tags			markets
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string					true	"User ID"
//	@Success		200		{object}	map[string]interface{}	"data object"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/markets/my/{user_id}  [get]
func getMyPostHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("user_id")
		if postID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}
		repo := marketrepository.NewMarketStore(db)
		data, err := repo.GetAllPost()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
