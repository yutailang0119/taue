package jobs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Yu-taro/taue/taue/models"
)

// GetContributs from API
func GetContributs() {

	users := loadUsersFromJSON()

	for _, user := range users {
		// fmt.Printf("%d : %s\n", user.ID, user.Name)
		getGitHubContributs(user)
	}

}

func getGitHubContributs(user models.User) {
	value := url.Values{}
	value.Add("access_token", user.GitHubToken)

	const baseURL = "https://api.github.com"
	urlString := baseURL + "/users/" + user.GitHubName + "/events"

	resp, err := http.Get(urlString + "?" + value.Encode())
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("%s", body)
		return
	}

	var githubEvents []models.GitHubEvent
	err = json.Unmarshal(body, &githubEvents)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user.GitHubEvents = githubEvents
	// for _, githubEvent := range user.GitHubEvents {
	// 	fmt.Printf("%s : %s\n", githubEvent.ID, githubEvent.Type)
	// }

}
