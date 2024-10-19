package markethandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterMarketRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/markets")
	{
		eg.POST("", createPostHandler(db))
		eg.GET("", getAllPostHandler(db))
		eg.GET("/:post_id", getPostDetailsHandler(db))
	}
}
