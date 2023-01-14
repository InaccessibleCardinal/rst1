package ctrl

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	testSuccessService    = MockSuccessService{}
	testFailureService    = MockFailureService{}
	testController        = NewUserController(&testSuccessService)
	testFailureController = NewUserController(&testFailureService)
)

func verifyBody(t *testing.T, actual []byte, expected []byte) {
	actualStr := string(actual)
	expectedStr := string(expected)
	if strings.Compare(actualStr, expectedStr) != 1 {
		t.Errorf("expected %s but got %s", expectedStr, actualStr)
	}
}

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
	resp := rec.Result()
	body, _ := io.ReadAll(resp.Body)
	j, _ := json.Marshal(mockUsers)
	verifyBody(t, body, j)
}

func TestGetUserFailure(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(testFailureController.HandleGetUser)
	handler.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
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
			status, http.StatusBadRequest)
	}
}
