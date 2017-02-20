package job

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type UserJson struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SlackName string `json:"slackName"`
	Github    string `json:"github"`
	Gitlab    string `json:"gitlab"`
}

func LoadUsersFromJson() []UserJson {

	file, err := ioutil.ReadFile("./resource/Users.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var users []UserJson

	json.Unmarshal(file, &users)

	return users

}
