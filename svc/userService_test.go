package svc

import (
	"rst1/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockSuccessUserRepository = mocks.MockSuccessUserRepository{}
	mockFailureUserRepository = mocks.MockFailureUserRepository{}
	testService               = NewUserService(&mockSuccessUserRepository)
	testFailureService        = NewUserService(&mockFailureUserRepository)
)

func TestFindAllSuccess(t *testing.T) {
	uResp := testService.FindAll()
	assert.Nil(t, uResp.Error)
	assert.Equal(t, uResp.IsOk, true)
	assert.Equal(t, uResp.Value[0].Firstname, "testFn")
}

func TestFindAllFailure(t *testing.T) {
	uResp := testFailureService.FindAll()
	assert.NotNil(t, uResp.Error)
	assert.Equal(t, uResp.IsOk, false)
}

func TestFindByIdSuccess(t *testing.T) {
	uResp := testService.FindById(0)
	assert.Nil(t, uResp.Error)
	assert.Equal(t, uResp.IsOk, true)
	assert.Equal(t, uResp.Value.Firstname, "testFn")
}

func TestFindByIdFailure(t *testing.T) {
	uResp := testService.FindById(1)
	assert.NotNil(t, uResp.Error)
	assert.Equal(t, uResp.IsOk, false)
}
