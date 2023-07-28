package master_author

import (
	"project-rest-api/entities"

	"gorm.io/gorm"
)

type Repository interface {
	GetMasterAuthor() ([]entities.MasterAuthor, error)
	StoreMasterAuthor(masterAuthor entities.MasterAuthor) (entities.MasterAuthor, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetMasterAuthor() ([]entities.MasterAuthor, error) {
	var ma []entities.MasterAuthor

	// SELECT * FROM master_authors
	if err := r.db.Find(&ma).Error; err != nil {
		return nil, err
	}

	return ma, nil
}

func (r *repository) StoreMasterAuthor(masterAuthor entities.MasterAuthor) (entities.MasterAuthor, error) {
	err := r.db.Create(&masterAuthor).Error

	if err != nil {
		return masterAuthor, err
	}

	return masterAuthor, nil
}
