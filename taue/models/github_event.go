package models

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
