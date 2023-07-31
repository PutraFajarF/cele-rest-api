package main

import (
	"log"
	"project-rest-api/auth"
	"project-rest-api/config"
	"project-rest-api/handler"
	"project-rest-api/master_author"
	"project-rest-api/master_book"
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
	masterAuthorRepository := master_author.NewRepository(db)
	masterBookRepository := master_book.NewRepository(db)

	// call service
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	masterAuthorService := master_author.NewService(masterAuthorRepository)
	masterBookService := master_book.NewService(masterBookRepository)

	// call handler
	userHandler := handler.NewUserHandler(userService, authService)
	masterAuthorHandler := handler.NewMasterAuthorHandler(masterAuthorService)
	masterBookHandler := handler.NewMasterBookHandler(masterBookService)

	// gin router
	router := gin.Default()

	// api versioning
	userApi := router.Group("/api/v1/user")
	masterApi := router.Group("/api/v1/master")

	routes.UserRoutes(userApi, userHandler)
	routes.MasterAuthorRoutes(masterApi, masterAuthorHandler, db, authService, userService)
	routes.MasterBookRoutes(masterApi, masterBookHandler)

	router.Run()
}
