package svc

import (
	"encoding/json"
	"rst1/ht"
	"rst1/models"
	"sync"
)

const remoteApi = "https://jsonplaceholder.typicode.com/users/1"

type RemoteService interface{}

type RemoteUserService struct {
	makeRequest ht.HtFunc
	dbService   UserService
}

func NewRemoteUserService(makeRequest ht.HtFunc, service UserService) *RemoteUserService {
	return &RemoteUserService{
		makeRequest: makeRequest,
		dbService:   service,
	}
}

func extractAddress(jsonMap map[string]interface{}) *models.Address {
	addr := jsonMap["address"].(map[string]interface{})
	city := addr["city"]
	street := addr["street"]
	suite := addr["suite"]
	zipcode := addr["zipcode"]
	return &models.Address{
		City:    city.(string),
		Street:  street.(string),
		Suite:   suite.(string),
		Zipcode: zipcode.(string)}
}

func getAddress() ServiceResponse[*models.Address] {
	jsonMap := make(map[string]interface{})
	resp, err := ht.MakeHttpRequest("GET", remoteApi, map[string]any{})
	if err != nil {
		return ServiceResponse[*models.Address]{IsOk: false, Value: nil, Error: err}
	}
	jerr := json.Unmarshal([]byte(resp), &jsonMap)
	if jerr != nil {
		return ServiceResponse[*models.Address]{IsOk: false, Value: nil, Error: jerr}
	}
	addr := extractAddress(jsonMap)
	return ServiceResponse[*models.Address]{IsOk: true, Value: addr, Error: nil}
}

func (ru *RemoteUserService) GetExpandedUser(id int) ServiceResponse[*models.UserExpanded] {
	var (
		wg              sync.WaitGroup
		addressResponse ServiceResponse[*models.Address]
		userResponse    ServiceResponse[models.User]
	)
	wg.Add(2)
	go func() {
		defer wg.Done()
		userResponse = ru.dbService.FindById(id)
	}()

	go func() {
		defer wg.Done()
		addressResponse = getAddress()
	}()
	wg.Wait()
	value := &models.UserExpanded{
		User:    userResponse.Value,
		Address: *addressResponse.Value,
	}
	return ServiceResponse[*models.UserExpanded]{IsOk: true, Value: value, Error: nil}
}
