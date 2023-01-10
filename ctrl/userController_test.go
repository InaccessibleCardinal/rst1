package ctrl

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testSuccessService    = MockSuccessService{}
	testFailureService    = MockFailureService{}
	testController        = NewUserController(&testSuccessService)
	testFailureController = NewUserController(&testFailureService)
)

func TestGetUsersSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(testController.HandleGetUsers)
	handler.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetUserFailure(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(testFailureController.HandleGetUser)
	handler.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestGetUserBadRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(testController.HandleGetUser)
	handler.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
