package json_dealing

import (
	"encoding/json"
)

type User struct {
	Name        string  `json:"name"`
	Age         int     `json:"age"`
	Nationality string  `json:"nationality"`
	Wallet      float32 `json:"wallet"`
	Currency    string  `json:"currency"`
	// CreatedAt   time.Time `json:"createdAt"`
}

func ParseUsers(response string) ([]User, error) {
	var users []User
	if err := json.Unmarshal([]byte(response), &users); err != nil {
		return nil, err
	}

	return users, nil
}
