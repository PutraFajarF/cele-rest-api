package master_book

import (
	"project-rest-api/entities"

	"gorm.io/gorm"
)

type Service interface {
	GetBooks() ([]entities.MasterBook, error)
	CreateMasterBook(input MasterBookInput) (entities.MasterBook, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
