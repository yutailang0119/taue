package models

import (
	"fmt"
	"time"
)

// User defind a user from json
type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	SlackName    string `json:"slackName"`
	GitHubName   string `json:"githubName"`
	GitHubToken  string `json:"githubToken"`
	GitLabID     int    `json:"gitlabId"`
	GitLabToken  string `json:"gitlabToken"`
	GitHubEvents []GitHubEvent
	GitLabEvents []GitLabEvent
}

// TodayGitHubContributs is count today activity on GitHub
func (user User) TodayGitHubContributs() int {
	var count int
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	fmt.Println(yesterday)

	for _, githubEvent := range user.GitHubEvents {
		fmt.Println(githubEvent.createdAt())
		count++
	}

	return count
}

// TodayGitLabContributs is count today activity on GitLab
func (user User) TodayGitLabContributs() int {
	var count int
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	fmt.Println(yesterday)

	for _, gitlabEvent := range user.GitLabEvents {
		fmt.Println(gitlabEvent.createdAt())
		count++
	}

	return count
}

// TodayContributs is count today activity on GitHub and GitLab
func (user User) TodayContributs() int {
	return user.TodayGitHubContributs() + user.TodayGitLabContributs()
}
