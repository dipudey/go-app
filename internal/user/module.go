package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(routeGroup *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	// Register routes
	SetupRoutes(routeGroup, handler)
}
