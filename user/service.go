package user

import (
	"errors"
	"project-rest-api/entities"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (entities.User, error)
	// LoginUser(input LoginInput) (entities.User, error)
	GetUserByID(ID int) (entities.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (entities.User, error) {
	var user entities.User
	user.Name = input.Name
	user.Email = input.Email
	user.NoHandphone = input.NoHandphone

	checkUserEmail, err := s.repository.FindUserByEmail(user.Email)
	if err != nil {
		return checkUserEmail, err
	}

	if checkUserEmail.ID != 0 {
		return checkUserEmail, errors.New("email ini sudah terdaftar")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.CreateUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) GetUserByID(ID int) (entities.User, error) {
	user, err := s.repository.FindUserByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("data user tidak ditemukan berdasarkan ID")
	}

	return user, nil
}