package user

import (
	"github.com/dipudey/go-app/internal/auth"
	"github.com/gin-gonic/gin"
)

// SetupRoutes registers user module routes.
// It does NOT know about repository or service — only receives a handler.
func SetupRoutes(rg *gin.RouterGroup, h *Handler) {
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/login", h.Login)
	}

	users := rg.Group("/users")
	users.Use(auth.JWTMiddleware())
	{
		users.GET("/", h.GetAll)
		users.POST("/create", h.Create)
	}
}
