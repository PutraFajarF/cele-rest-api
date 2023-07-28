package user

import (
	"project-rest-api/entities"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(entities.User) (entities.User, error)
	FindUserByEmail(email string) (entities.User, error)
	FindUserByID(ID int) (entities.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(user entities.User) (entities.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindUserByEmail(email string) (entities.User, error) {
	var user entities.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindUserByID(ID int) (entities.User, error) {
	var user entities.User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
