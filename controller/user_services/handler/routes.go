package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterUserRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/users")
	{
		eg.GET("/hello", sayHello)
	}
}

// @BasePath		/api/v1
// @Summary		hello
// @Description	hello
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{object}	string	"say hi"
// @Router			/users/hello [get]
func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "hello from Phu Phuoc"})
}
