package markethandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterMarketRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/markets")
	{
		eg.GET("", getAllPostHandler(db))
		eg.GET("/my/:user_id", getMyPostHandler(db))
		eg.POST("", createPostHandler(db))
		eg.GET("/:post_id", getPostDetailsHandler(db))
		eg.PUT("/:post_id", updatePostDetailHandler(db))
		eg.DELETE("/:post_id", deletePostDetailsHandler(db))
	}
}
