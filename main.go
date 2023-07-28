package main

import (
	"log"
	"project-rest-api/auth"
	"project-rest-api/config"
	"project-rest-api/handler"
	"project-rest-api/routes"
	"project-rest-api/user"

	"github.com/gin-gonic/gin"
)

// flow code = main.go => routes => handler => service => repository
func main() {
	config.LoadAppConfig()
	db, err := config.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("Database sukses terkoneksi")

	err = config.Migrate()

	if err != nil {
		log.Fatal(err.Error())
	}

	// Call repository
	userRepository := user.NewRepository(db)

	// call service
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	// call handler
	userHandler := handler.NewUserHandler(userService, authService)

	// gin router
	router := gin.Default()

	// api versioning
	userApi := router.Group("/api/v1/user")

	routes.UserRoutes(userApi, userHandler)

	router.Run()
}
