package routes

import (
	"project-rest-api/auth"
	"project-rest-api/handler"
	"project-rest-api/middleware"
	"project-rest-api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MasterAuthorRoutes(api *gin.RouterGroup, handler *handler.MasterAuthorHandler, db *gorm.DB, authService auth.Service, userService user.Service) {
	api.GET("/author", handler.Get)
	api.POST("/author/store", middleware.AuthMiddleware(authService, userService), handler.Create)
}
