package models

// User defind a user from json
type User struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	SlackName        string `json:"slackName"`
	GitHubName       string `json:"githubName"`
	GitHubToken      string `json:"githubToken"`
	GitLabName       string `json:"gitlabName"`
	GitLabToken      string `json:"gitlabToken"`
	GitHubEvents     []GitHubEvent
	GitHubContributs int
	GitLabContributs int
}
