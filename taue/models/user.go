package models

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

// todayGitHubContributs is count today activity on GitHub
func (user User) todayGitHubContributs() int {
	var count int

	for _, githubEvent := range user.GitHubEvents {
		if githubEvent.isTodayContribute() {
			count++
		}
	}

	return count
}

// todayGitLabContributs is count today activity on GitLab
func (user User) todayGitLabContributs() int {
	var count int

	for _, gitlabEvent := range user.GitLabEvents {
		if gitlabEvent.isTodayContribute() {
			count++
		}
	}

	return count
}

// TodayContributs is count today activity on GitHub and GitLab
func (user User) TodayContributs() int {
	return user.todayGitHubContributs() + user.todayGitLabContributs()
}
