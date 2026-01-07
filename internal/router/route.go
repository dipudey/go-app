package router

import (
	_ "github.com/dipudey/go-app/docs"
	"github.com/dipudey/go-app/internal/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) *gin.Engine {
	//router := gin.Default()
	router := gin.New()
	router.Use(gin.Recovery())
	api := router.Group("/api")

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello go app")
	})

	router.GET("/api/ping", PingHandler)

	// Register user module routes
	user.Register(api, db)

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

// PingHandler PingExample godoc
// @Summary Ping the server
// @Description Returns "pong" to test API
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
