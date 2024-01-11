package json_dealing

import (
	"testing"
)

func TestParseUsersFromString(t *testing.T) {
	response := `{"users": [{"name": "Louis","age": "19","nationality": "German","wallet": "3526.3","currency": "EUR","createdAt": "2021-10-18T11:08:47.577Z"}]}`
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
		// CreatedAt:   parsed[0].CreatedAt,
	}

	if parsed[0] != expected {
		t.Fatalf("First user is different from expected one")
	}
}
