package models

// SlackParameters is Slack WebHook parameters model
type SlackParameters struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconURL   string `json:"icon_url"`
	Channel   string `json:"channel"`
	LinkNames int    `json:"link_names"`
}
