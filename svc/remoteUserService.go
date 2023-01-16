package svc

import (
	"encoding/json"
	"fmt"
	"rst1/ht"
	"rst1/models"
	"strconv"
	"sync"

	"github.com/patrickmn/go-cache"
)

const remoteApi = "https://jsonplaceholder.typicode.com/users/"

type RemoteService interface{}

type RemoteUserService struct {
	makeRequest ht.HtFunc
	dbService   UserService
	cache       *cache.Cache
}

func NewRemoteUserService(makeRequest ht.HtFunc, service UserService, cache *cache.Cache) *RemoteUserService {
	return &RemoteUserService{
		makeRequest: makeRequest,
		dbService:   service,
		cache:       cache,
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

func (ru *RemoteUserService) checkCache(id string) (*models.Address, bool) {
	fmt.Printf("hitting cache for id %s\n", id)
	addr, found := ru.cache.Get(id)
	if found {
		return addr.(*models.Address), true
	}
	return nil, false
}

func (ru *RemoteUserService) getAddress(id int) ServiceResponse[*models.Address] {
	jsonMap := make(map[string]interface{})

	addr, found := ru.checkCache(strconv.Itoa(id))
	if found {
		fmt.Printf("using cache for id %d\n", id)
		return ServiceResponse[*models.Address]{IsOk: false, Value: addr, Error: nil}
	}

	resp, err := ht.MakeHttpRequest("GET", remoteApi+strconv.Itoa(id), map[string]any{})
	if err != nil {
		return ServiceResponse[*models.Address]{IsOk: false, Value: nil, Error: err}
	}
	jerr := json.Unmarshal([]byte(resp), &jsonMap)
	if jerr != nil {
		return ServiceResponse[*models.Address]{IsOk: false, Value: nil, Error: jerr}
	}
	address := extractAddress(jsonMap)
	ru.cache.Add(strconv.Itoa(id), address, cache.DefaultExpiration)
	return ServiceResponse[*models.Address]{IsOk: true, Value: address, Error: nil}
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
		addressResponse = ru.getAddress(id)
	}()
	wg.Wait()
	value := &models.UserExpanded{
		User:    userResponse.Value,
		Address: *addressResponse.Value,
	}
	return ServiceResponse[*models.UserExpanded]{IsOk: true, Value: value, Error: nil}
}
