package commenthandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterCommentRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/post/:post_id/comment")
	{
		eg.GET("", getAllCommentHandler(db))
		eg.POST("", addCommentHandler(db))
		eg.DELETE("/:comment_id", deleteCommentHandler(db))
	}
}
