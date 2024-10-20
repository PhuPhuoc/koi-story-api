package consulthandler

import (
	"net/http"

	consultmodel "github.com/PhuPhuoc/koi-story-api/controller/consult_services/model"
	consultrepository "github.com/PhuPhuoc/koi-story-api/controller/consult_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// CreateNewConsultPost godoc
//
//	@Summary		Create new consult's post
//	@Description	Create new consult's post
//	@Tags			consults
//	@Accept			json
//	@Produce		json
//	@Param			user	body		consultmodel.CreateConsultInfo	true	"details of new post of consult"
//	@Success		201		{object}	map[string]interface{}		"message success"
//	@Failure		400		{object}	error						"Bad request error"
//	@Router			/consults [post]
func createNewConsultPostHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req consultmodel.CreateConsultInfo

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		repo := consultrepository.NewConsultStore(db)

		err := repo.CreateNewConsult(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "created new post successfully"})
	}
}
