package api

import (
	"fmt"

	markethandler "github.com/PhuPhuoc/koi-story-api/controller/market_services/handler"
	userhandler "github.com/PhuPhuoc/koi-story-api/controller/user_services/handler"
	docs "github.com/PhuPhuoc/koi-story-api/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type server struct {
	address         string
	db              *sqlx.DB
	host_production string
}

func InitServer(addr, host string, db *sqlx.DB) *server {
	return &server{
		address:         addr,
		db:              db,
		host_production: host,
	}
}

func (sv *server) RunApp() error {
	docs.SwaggerInfo.BasePath = "/api/v1"
	router := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/swagger/*any"),
		gin.Recovery(),
		cors.New(config),
	)

	v1 := router.Group("/api/v1")
	sv.registerRoutes(v1)
	sv.runLog()
	return router.Run(sv.address)
}

func (sv *server) runLog() {
	fmt.Printf("\nFor dev: http://localhost%s/swagger/index.html\n", sv.address)
	fmt.Printf("\nFor production: http://%s%s/swagger/index.html\n", sv.host_production, sv.address)
}

func (sv *server) registerRoutes(v1 *gin.RouterGroup) {
	userhandler.RegisterUserRoutes(v1, sv.db)
	markethandler.RegisterMarketRoutes(v1, sv.db)
}
