package jobs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/Yu-taro/taue/taue/models"
)

func postSlack(users []models.User) {
	webHookURL := "https://hooks.slack.com/services/" + os.Getenv("SLACK_WEBHOOK_ENDPOINT")

	var text string
	for _, user := range users {
		text = text + "@" + user.SlackName + " " + user.Contributes() + "\n"
	}
	log.Print(text)

	parameters := models.SlackParameters{
		Text:      text,
		Username:  "taue",
		IconEmoji: ":seedling:",
		IconURL:   "",
		Channel:   "",
		LinkNames: 1,
	}

	params, _ := json.Marshal(parameters)

	value := url.Values{"payload": {string(params)}}
	resp, err := http.PostForm(webHookURL, value)

	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body))

}
