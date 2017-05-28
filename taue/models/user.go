package models

import (
	//"database/sql"
)
import "database/sql"

// User defind a user from json
type User struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	SlackName    string `db:"slack_name"`
	GitHubName   sql.NullString `db:"github_name"`
	GitHubToken  sql.NullString `db:"github_token"`
	GitLabID     sql.NullInt64    `db:"gitlab_id"`
	GitLabToken  sql.NullString `db:"gitlab_token"`
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
