package jobs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/yutailang0119/taue/taue/models"
)

// loadUsersFromJSON load users profile from Users.json
func loadUsersFromJSON() (users []models.User) {

	file, err := ioutil.ReadFile("taue/resources/Users.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	json.Unmarshal(file, &users)

	return users

}
