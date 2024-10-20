package consulthandler

import (
	"net/http"

	consultrepository "github.com/PhuPhuoc/koi-story-api/controller/consult_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// GetAllConsultPost godoc
//
//	@Summary		Get all consult's post
//	@Description	Get all consult's post
//	@Tags			consults
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"data object"
//	@Failure		400	{object}	error					"Bad request error"
//	@Router			/consults [get]
func getAllConsultPostHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo := consultrepository.NewConsultStore(db)
		data, err := repo.GetListConsult()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
