package user

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes registers user module routes.
// It does NOT know about repository or service — only receives a handler.
func SetupRoutes(rg *gin.RouterGroup, h *Handler) {
	users := rg.Group("/users")
	{
		users.GET("/", h.GetAll)
	}
}
