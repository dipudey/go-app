package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) *gin.Engine {
	//router := gin.Default()
	router := gin.New()
	router.Use(gin.Recovery())
	//api := router.Group("/api")

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello go app")
	})

	return router
}
