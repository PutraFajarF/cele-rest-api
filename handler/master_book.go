package handler

import (
	"net/http"
	"project-rest-api/helper"
	"project-rest-api/master_book"

	"github.com/gin-gonic/gin"
)

type MasterBookHandler struct {
	masterBookService master_book.Service
}

func NewMasterBookHandler(masterBookService master_book.Service) *MasterBookHandler {
	return &MasterBookHandler{masterBookService}
}

func (mb *MasterBookHandler) Get(c *gin.Context) {
	masterBook, err := mb.masterBookService.GetBooks()

	if err != nil {
		response := helper.JsonResponse("Gagal mendapatkan data buku", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("data buku berhasil diambil", http.StatusOK, "success", masterBook)

	c.JSON(http.StatusOK, response)
}

func (mb *MasterBookHandler) Create(c *gin.Context) {
	var input master_book.MasterBookInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal menyimpan data buku", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newMasterBook, err := mb.masterBookService.CreateMasterBook(input)

	if err != nil {
		response := helper.JsonResponse("Gagal menyimpan data buku", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := master_book.FormatMasterBook(newMasterBook)
	response := helper.JsonResponse("Penyimpanan data buku berhasil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}
