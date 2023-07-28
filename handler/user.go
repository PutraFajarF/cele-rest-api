package handler

import (
	"net/http"
	"project-rest-api/auth"
	"project-rest-api/helper"
	"project-rest-api/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *UserHandler {
	return &UserHandler{userService, authService}
}

func (u *UserHandler) RegisterUser(c *gin.Context) {
	// Tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct diatas kita passing sebagai parameter service
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal register akun", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := u.userService.RegisterUser(input)
	if err != nil {
		response := helper.JsonResponse("Gagal register akun", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := u.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.JsonResponse("Gagal register akun", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := user.FormatUser(newUser, token)
	response := helper.JsonResponse("Registrasi akun berhasil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}
