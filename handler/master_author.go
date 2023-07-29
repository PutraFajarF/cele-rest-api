package handler

import (
	"net/http"
	"project-rest-api/helper"
	"project-rest-api/master_author"

	"github.com/gin-gonic/gin"
)

type MasterAuthorHandler struct {
	masterAuthorService master_author.Service
}

func NewMasterAuthorHandler(masterAuthorService master_author.Service) *MasterAuthorHandler {
	return &MasterAuthorHandler{masterAuthorService}
}

func (mb *MasterAuthorHandler) Get(c *gin.Context) {
	masterAuthor, err := mb.masterAuthorService.GetAuthors()

	if err != nil {
		response := helper.JsonResponse("Gagal mendapatkan data author", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("data author berhasil diambil", http.StatusOK, "success", masterAuthor)

	c.JSON(http.StatusOK, response)
}

func (mb *MasterAuthorHandler) Create(c *gin.Context) {
	var input master_author.MasterAuthorInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal menyimpan data author", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newMasterAuthor, err := mb.masterAuthorService.CreateMasterAuthor(input)

	if err != nil {
		response := helper.JsonResponse("Gagal menyimpan data author", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := master_author.FormatMasterAuthor(newMasterAuthor)
	response := helper.JsonResponse("Penyimpanan data author berhasil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}
