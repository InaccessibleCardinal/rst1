package svc

import (
	"encoding/json"
	"testing"
)

func TestExtractAddress(t *testing.T) {
	testAddr := `{
		"address": {
			"city": "plano",
			"street": "1 elm st",
			"suite": "3",
			"zipcode": "75555"
		}
	}`
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
