package models

import "time"

// GitLabEvent is https://gitlab.com/api/v3/users/:id/events/
type GitLabEvent struct {
	Title          string      `json:"title"`
	ProjectID      int         `json:"project_id"`
	ActionName     string      `json:"action_name"`
	TargetID       int         `json:"target_id"`
	TargetType     string      `json:"target_type"`
	TargetTitle    string      `json:"target_title"`
	AuthorID       int         `json:"author_id"`
	AuthorUserName string      `json:"author_username"`
	Autor          gitlabAutor `json:"author"`
	CreatedAt      string      `json:"created_at"`
}

type gitlabAutor struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	UserName  string `json:"username"`
	Status    string `json:"status"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

func (gitlabEvent GitLabEvent) createdAt() time.Time {
	const gitlabTimeLayout = "2006-01-02T15:04:05.000Z"
	t, _ := time.Parse(gitlabTimeLayout, gitlabEvent.CreatedAt)
	return t
}
