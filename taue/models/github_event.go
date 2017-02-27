package models

import (
	"time"
)

// GitHubEvent is https://api.github.com/users/:username/event
type GitHubEvent struct {
	ID        string      `json:"id"`
	Type      string      `json:"type"`
	Actor     githubActor `json:"actor"`
	Repo      githubRepo  `json:"repo"`
	CreatedAt string      `json:"created_at"`
}

type githubActor struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

type githubRepo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (githubEvent GitHubEvent) createdAt() time.Time {
	const githubTimeLayout = "2006-01-02T15:04:05Z"
	t, _ := time.Parse(githubTimeLayout, githubEvent.CreatedAt)
	return t
}

func (githubEvent GitHubEvent) isTodayContribute() bool {
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	thisCreatedAt := githubEvent.createdAt()
	return thisCreatedAt.Year() == yesterday.Year() && thisCreatedAt.Month() == yesterday.Month() && thisCreatedAt.Day() == yesterday.Day()
}
