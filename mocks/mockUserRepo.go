package mocks

import (
	"errors"
	"rst1/models"
)

var testUser = models.User{
	Firstname: "testFn",
	Lastname:  "testLn",
	Email:     "test@test.com",
	Password:  "testpass",
	Id:        0,
}

type MockSuccessUserRepository struct{}
type MockFailureUserRepository struct{}

func (m *MockSuccessUserRepository) FindAll() ([]models.User, error) {
	testResp := []models.User{
		testUser,
	}
	return testResp, nil
}

func (m *MockFailureUserRepository) FindAll() ([]models.User, error) {
	return nil, errors.New("oops")
}

func (m *MockSuccessUserRepository) FindById(id int) (models.User, error) {
	if id == 0 {
		return testUser, nil
	}
	return models.User{}, errors.New("oops")
}

func (m *MockFailureUserRepository) FindById(id int) (models.User, error) {
	if id == 0 {
		return testUser, nil
	}
	return models.User{}, errors.New("oops")
}

func (m *MockSuccessUserRepository) Update(u *models.User) (int64, error) {

	return 0, errors.New("oops")
}

func (m *MockFailureUserRepository) Update(u *models.User) (int64, error) {

	return 0, errors.New("oops")
}
