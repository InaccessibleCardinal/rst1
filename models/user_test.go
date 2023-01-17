package models

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUserToJson(t *testing.T) {
	testUser := User{
		Firstname: "fn",
		Lastname:  "ln",
		Email:     "fnlf@site.com",
		Password:  "pass",
		Id:        1,
	}
	res, err := json.Marshal(testUser)
	if err != nil {
		t.Error("failed to marshal user struct")
	}
	if !strings.Contains(string(res), `"firstname":"fn"`) {
		t.Errorf("failed to properly serialze %s", res)
	}
}
