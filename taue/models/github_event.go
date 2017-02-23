package models

type GitHubEvent struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Actor     actor  `json:"actor"`
	Repo      repo   `json:"repo"`
	CreatedAt string `json:"created_at"`
}

type actor struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

type repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
