package svc

import (
	"rst1/models"
)

type UserService struct {
	repository models.UserRepository
}

func NewUserService(repository models.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (userService *UserService) FindAll() ServiceResponse[[]models.User] {
	users, err := userService.repository.FindAll()
	if err != nil {
		return ServiceResponse[[]models.User]{IsOk: false, Value: nil, Error: err}
	}
	return ServiceResponse[[]models.User]{IsOk: true, Value: users, Error: nil}
}

func (userService *UserService) FindById(id int) ServiceResponse[models.User] {
	user, err := userService.repository.FindById(id)
	if err != nil {
		println("no user found for id: ", id)
		return ServiceResponse[models.User]{IsOk: false, Value: user, Error: err}
	}
	return ServiceResponse[models.User]{IsOk: true, Value: user, Error: nil}
}

func (userService *UserService) Update(user *models.User) ServiceResponse[*models.User] {
	_, err := userService.repository.Update(user)
	if err != nil {
		return ServiceResponse[*models.User]{IsOk: false, Value: nil, Error: err}
	}
	return ServiceResponse[*models.User]{IsOk: true, Value: user, Error: nil}
}
