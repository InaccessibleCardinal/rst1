package svc

import (
	"encoding/json"
	"rst1/models"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	mockUser = models.User{
		Firstname: "fn",
		Lastname:  "ln",
		Email:     "fnln@site.com",
		Password:  "pw",
		Id:        1,
	}
	mockUserStr = `{
		"firstname": "fn",
		"lastname": "ln",
		"email": "fnln@site.com",
		"password": "pw",
		"id": 1
	}`
	testAddr = `{
		"address": {
			"city": "plano",
			"street": "1 elm st",
			"suite": "3",
			"zipcode": "75555"
		}
	}`
)

func TestExtractAddress(t *testing.T) {
	jsonAddr := &map[string]any{}
	err := json.Unmarshal([]byte(testAddr), jsonAddr)
	if err != nil {
		t.Errorf("error parsing: %s", err.Error())
	}
	result := extractAddress(*jsonAddr)
	if result.City != "plano" {
		t.Errorf("expected `plano`, got %s", result.City)
	}
	if result.Street != "1 elm st" {
		t.Errorf("expected `1 elm st`, got %s", result.City)
	}
}

type MockService struct{}

func (m *MockService) FindById(id int) ServiceResponse[models.User] {
	return ServiceResponse[models.User]{IsOk: true, Value: mockUser, Error: nil}
}

func (m *MockService) FindAll() ServiceResponse[[]models.User] {
	mockUsers := make([]models.User, 10)
	return ServiceResponse[[]models.User]{IsOk: true, Value: mockUsers, Error: nil}
}

func (m *MockService) Update(u *models.User) ServiceResponse[*models.User] {
	return ServiceResponse[*models.User]{IsOk: true, Value: u, Error: nil}
}

func makeReq(method string, url string, body map[string]any) (string, error) {
	return testAddr, nil
}
func TestGetAddress(t *testing.T) {
	testCache := cache.New(1*time.Second, 10*time.Second)
	testService := &MockService{}
	r := NewRemoteUserService(makeReq, testService, testCache)
	goodResponse := r.getAddress(1)
	if !goodResponse.IsOk {
		t.Errorf("expected IsOk == true")
	}
	if goodResponse.Value.City != "plano" {
		t.Errorf("expected plano got %s", goodResponse.Value.City)
	}
}
