package consulthandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterConsultRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/consults")
	{
		eg.GET("", getAllConsultPostHandler(db))
		eg.GET("/my/:user_id", getMyConsultPostHandler(db))
		eg.POST("", createNewConsultPostHandler(db))
		eg.PUT("/:post_id", updatePostDetailHandler(db))
		eg.DELETE("/:post_id", deletePostDetailsHandler(db))
	}
}
