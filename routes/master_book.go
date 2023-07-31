package routes

import (
	"project-rest-api/handler"

	"github.com/gin-gonic/gin"
)

func MasterBookRoutes(api *gin.RouterGroup, handler *handler.MasterBookHandler) {
	api.GET("/book", handler.Get)
	api.POST("/book/store", handler.Create)
}
