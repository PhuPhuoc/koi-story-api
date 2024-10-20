package consulthandler

import (
	"net/http"

	consultrepository "github.com/PhuPhuoc/koi-story-api/controller/consult_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// GetMyConsultPost godoc
//
//	@Summary		Get all my consult's post
//	@Description	Get all my consult's post
//	@Tags			consults
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string					true	"User ID"
//	@Success		200		{object}	map[string]interface{}	"data object"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/consults/my/{user_id}  [get]
func getMyConsultPostHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		if user_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}
		repo := consultrepository.NewConsultStore(db)
		data, err := repo.GetMyListConsult(user_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
