package json_dealing

import (
	"encoding/json"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name        string  `json:"name"`
	Age         int     `json:"age"`
	Nationality string  `json:"nationality"`
	Wallet      float32 `json:"wallet"`
	Currency    string  `json:"currency"`
}

func ParseUser(data []byte) (User, error) {
	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		return User{}, err
	}

	return user, nil
}

func ParseUsers(data []byte) ([]User, error) {
	var users Users
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users.Users, nil
}
