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

func (u *UserHandler) Login(c *gin.Context) {
	// user masukan input (email & password)
	// input ditangkap handler
	// mapping input user ke input struct
	// input struct passing ke service
	// service mencari dengan bantuan repository user dengan email x
	// cocokan password
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Akun gagal login", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginUser, err := u.userService.LoginUser(input)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}

		response := helper.JsonResponse("Akun gagal login", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := u.authService.GenerateToken(loginUser.ID)
	if err != nil {
		response := helper.JsonResponse("Akun gagal login", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := user.FormatUser(loginUser, token)
	response := helper.JsonResponse("Login sukses", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}
