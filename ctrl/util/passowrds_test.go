package util

import "testing"

func TestHashAndCompare(t *testing.T) {
	testPass := "lolwut"
	hashed, err := HashPassword(testPass)
	if err != nil {
		t.Error("somethign went wrong")
	}
	matches := CheckPasswordHash(testPass, hashed)
	if !matches {
		t.Error("hash is incorrect for test password")
	}
}
