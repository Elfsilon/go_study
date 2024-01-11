package json_dealing

import (
	"testing"
)

func TestParseUserFromString(t *testing.T) {
	response := `{"name": "Louis","age": 19,"nationality": "German","wallet": 3526.3,"currency": "EUR","createdAt": "2021-10-18T11:08:47.577Z"}`
	parsed, err := ParseUser(response)

	if err != nil {
		t.Fatalf("Error ocurred %s", err)
	}

	expected := User{
		Name:        "Louis",
		Age:         19,
		Nationality: "German",
		Wallet:      3526.3,
		Currency:    "EUR",
	}

	if parsed != expected {
		t.Fatalf("Parsed user is different from expected one")
	}
}

func TestParseUsersFromString(t *testing.T) {
	response := `{"users": [{"name": "Louis","age": 19,"nationality": "German","wallet": 3526.3,"currency": "EUR","createdAt": "2021-10-18T11:08:47.577Z"},{"name": "Louis","age": 19,"nationality": "German","wallet": 3526.3,"currency": "EUR","createdAt": "2021-10-18T11:08:47.577Z"}]}`
	parsed, err := ParseUsers(response)

	if err != nil {
		t.Fatalf("Error ocurred %s", err)
	}

	expected := User{
		Name:        "Louis",
		Age:         19,
		Nationality: "German",
		Wallet:      3526.3,
		Currency:    "EUR",
	}

	if parsed[0] != expected {
		t.Fatalf("First parsed user is different from expected one")
	}
}
