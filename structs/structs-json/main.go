package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type user struct {
	User        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"`
}

func main() {
	users := []user{
		{"kalevi", "1234", nil},
		{"god", "42", permissions{"admin": true}},
		{"jehu", "47", permissions{"user": true, "viewer": true}},
	}

	out, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}
