package ht

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMakeHttpRequestSuccess(t *testing.T) {
	expectedResponse := `{"problems": 99}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, expectedResponse)
	}))
	defer ts.Close()
	testUrl := ts.URL

	actualResponse, err := MakeHttpRequest("GET", testUrl, map[string]any{})
	if err != nil {
		t.Fatalf("test http call failed %s", err.Error())
	}
	if strings.Compare(actualResponse, expectedResponse) != 1 {
		t.Fatalf("expected %s but got %s", expectedResponse, actualResponse)
	}
}

func TestMakeHttpRequestFailure(t *testing.T) {
	badUrl := "http://localhost:9999"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "lol")
	}))
	defer ts.Close()
	_, err := MakeHttpRequest("GET", badUrl, map[string]any{})
	if !strings.Contains(err.Error(), "connection refused") {
		t.Fatalf("expected a connection error, got %s", err.Error())
	}
}
