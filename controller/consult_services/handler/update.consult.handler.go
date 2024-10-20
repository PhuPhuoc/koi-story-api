package consulthandler

import (
	"net/http"

	consultmodel "github.com/PhuPhuoc/koi-story-api/controller/consult_services/model"
	consultrepository "github.com/PhuPhuoc/koi-story-api/controller/consult_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// UpdatePostConsult godoc
//
//	@Summary		Update consult's post
//	@Description	Update consult's post
//	@Tags			consults
//	@Accept			json
//	@Produce		json
//	@Param			post_id	path		string							true	"Post ID"
//	@Param			user	body		consultmodel.UpdateConsultInfo	true	"details of current post of consult"
//	@Success		200		{object}	map[string]interface{}			"message success"
//	@Failure		400		{object}	error							"Bad request error"
//	@Router			/consults/{post_id} [put]
func updatePostDetailHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		postID := c.Param("post_id")
		if postID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
			return
		}

		var req consultmodel.UpdateConsultInfo
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		repo := consultrepository.NewConsultStore(db)

		err := repo.UpdateConsult(postID, req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "updated post successfully"})
	}
}
