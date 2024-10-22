package imagehandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterImageRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/post/:post_id/image")
	{
		eg.POST("", addImageHandler(db))
		eg.DELETE("/:image_id", deleteImageInPostMarketHandler(db))
	}
}
