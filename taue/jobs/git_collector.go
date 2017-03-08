package jobs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/Yu-taro/taue/taue/models"
)

// getContributes from API
func getContributes(users []models.User) (completeHandler []models.User) {

	ch := make(chan models.User, len(users))
	var wg sync.WaitGroup

	for _, user := range users {
		wg.Add(1)
		go getGitActivity(&wg, &user, ch)
	}

	wg.Wait()
	close(ch)

	for data := range ch {
		completeHandler = append(completeHandler, data)
	}

	return completeHandler

}

func getGitActivity(wg *sync.WaitGroup, user *models.User, ch chan models.User) {
	getGitHubContributes(user)
	getGitLabContributes(user)
	ch <- *user
	wg.Done()
}

func getGitHubContributes(user *models.User) {
	if user.GitHubName == "" {
		return
	}

	value := url.Values{}
	value.Add("per_page", "100")
	if user.GitHubToken != "" {
		value.Add("access_token", user.GitHubToken)
	}

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

}

func getGitLabContributes(user *models.User) {
	if user.GitLabID == 0 {
		return
	}

	value := url.Values{}
	value.Add("per_page", "100")
	if user.GitLabToken != "" {
		value.Add("private_token", user.GitLabToken)
	}

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

}
