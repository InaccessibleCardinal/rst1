package ctrl

import (
	"errors"
	"rst1/models"
	"rst1/svc"
)

var mockUsers = []models.User{
	{
		Firstname: "testFn",
		Lastname:  "testLn",
		Email:     "test@test.com",
		Password:  "testpass",
		Id:        0},
}

type MockSuccessService struct{}
type MockFailureService struct{}

func (m *MockSuccessService) FindAll() svc.ServiceResponse[[]models.User] {
	return svc.ServiceResponse[[]models.User]{IsOk: true, Value: mockUsers, Error: nil}
}

func (m *MockFailureService) FindAll() svc.ServiceResponse[[]models.User] {
	return svc.ServiceResponse[[]models.User]{IsOk: false, Value: []models.User{}, Error: errors.New("oops")}
}

func (m *MockSuccessService) FindById(id int) svc.ServiceResponse[models.User] {
	if id == 0 {
		return svc.ServiceResponse[models.User]{IsOk: true, Value: mockUsers[0], Error: nil}
	}
	return svc.ServiceResponse[models.User]{IsOk: false, Value: models.User{}, Error: errors.New("oops")}
}

func (m *MockFailureService) FindById(id int) svc.ServiceResponse[models.User] {
	return svc.ServiceResponse[models.User]{IsOk: false, Value: models.User{}, Error: errors.New("oops")}
}

func (m *MockSuccessService) Update(u *models.User) svc.ServiceResponse[*models.User] {
	return svc.ServiceResponse[*models.User]{IsOk: false, Value: &models.User{}, Error: errors.New("oops")}
}

func (m *MockFailureService) Update(u *models.User) svc.ServiceResponse[*models.User] {
	return svc.ServiceResponse[*models.User]{IsOk: false, Value: &models.User{}, Error: errors.New("oops")}
}
