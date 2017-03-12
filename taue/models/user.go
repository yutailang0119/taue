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

// todayGitHubContributesCount is count today activity on GitHub
func (user User) todayGitHubContributesCount() int {
	var count int

	for _, githubEvent := range user.GitHubEvents {
		if githubEvent.isTodayContribute() && githubEvent.isCountable() {
			count++
		}
	}

	return count
}

// todayGitLabContributesCount is count today activity on GitLab
func (user User) todayGitLabContributesCount() int {
	var count int

	for _, gitlabEvent := range user.GitLabEvents {
		if gitlabEvent.isTodayContribute() {
			count++
		}
	}

	return count
}

// TodayContributesCount is count today activity on GitHub and GitLab
func (user User) TodayContributesCount() int {
	return user.todayGitHubContributesCount() + user.todayGitLabContributesCount()
}
