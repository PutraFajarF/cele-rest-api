package routes

import (
	"project-rest-api/handler"

	"github.com/gin-gonic/gin"
)

func MasterAuthorRoutes(api *gin.RouterGroup, handler *handler.MasterAuthorHandler) {
	api.GET("/author", handler.Get)
	api.POST("/author/store", handler.Create)
}
