package models

// User defind a user from json
type User struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	SlackName        string `json:"slackName"`
	GitHubName       string `json:"githubName"`
	GitHubToken      string `json:"githubToken"`
	GitLabID         int    `json:"gitlabId"`
	GitLabToken      string `json:"gitlabToken"`
	GitHubEvents     []GitHubEvent
	GitLabEvents     []GitLabEvent
	GitHubContributs int
	GitLabContributs int
}
