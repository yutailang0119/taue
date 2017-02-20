package main

import (
	"fmt"

	"github.com/Yu-taro/taue/taue/job"
)

func main() {
	var users = job.LoadUsersFromJson()

	for _, user := range users {
		fmt.Printf("%d : %s\n", user.ID, user.Name)
	}
}
