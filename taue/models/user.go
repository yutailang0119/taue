package models

// User defind a user from json
type User struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	SlackName        string `json:"slackName"`
	Github           string `json:"github"`
	Gitlab           string `json:"gitlab"`
	GithubContributs int
	GitlabContributs int
}
