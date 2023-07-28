package routes

import (
	"project-rest-api/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, handler *handler.UserHandler) {
	api.POST("/register", handler.RegisterUser)

}
