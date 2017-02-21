package main

import (
	"fmt"

	"github.com/Yu-taro/taue/taue/jobs"
)

func main() {
	var users = jobs.LoadUsersFromJSON()

	for _, user := range users {
		fmt.Printf("%d : %s\n", user.ID, user.Name)
	}
}
