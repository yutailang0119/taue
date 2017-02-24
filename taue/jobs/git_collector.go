package jobs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Yu-taro/taue/taue/models"
)

// GetContributs from API
func GetContributs() {

	users := loadUsersFromJSON()

	for _, user := range users {
		// fmt.Printf("%d : %s\n", user.ID, user.Name)
		getGitHubContributs(user)
		getGitLabContributs(user)
	}

}

func getGitHubContributs(user models.User) {
	value := url.Values{}
	value.Add("access_token", user.GitHubToken)
	value.Add("pre_page", "100")

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
	for _, githubEvent := range user.GitHubEvents {
		fmt.Printf("%s : %s\n", githubEvent.Actor.Login, githubEvent.Type)
	}

}

func getGitLabContributs(user models.User) {
	value := url.Values{}
	value.Add("private_token", user.GitLabToken)
	value.Add("per_page", "100")

	const baseURL = "https://gitlab.com/api/v3"
	urlString := baseURL + "/users/" + strconv.Itoa(user.GitLabID) + "/events"

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

	var gitlabEvents []models.GitLabEvent

	err = json.Unmarshal(body, &gitlabEvents)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user.GitLabEvents = gitlabEvents
	for _, gitlabEvent := range user.GitLabEvents {
		fmt.Printf("%s : %s\n", gitlabEvent.Title, gitlabEvent.CreatedAt)
	}
}
