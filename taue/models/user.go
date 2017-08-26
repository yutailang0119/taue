package models

import (
	"github.com/yutailang0119/taue/taue/db"
)

type UserList []User

// User define a user from json
type User struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	SlackName       string `json:"slack_name"`
	GitHubName      string `json:"github_name"`
	GitHubToken     string `json:"github_token"`
	GitLabID        int64    `json:"gitlab_id"`
	GitLabToken     string `json:"gitlab_token"`
	ActionHistories []ActionHistory
}

func (v *User) FromRow(vdb *db.User) error {
	v.ID = vdb.ID
	v.Name = vdb.Name
	v.SlackName = vdb.SlackName
	if vdb.GitHubName.Valid {
		v.GitHubName = vdb.GitHubName.String
	}
	if vdb.GitHubToken.Valid {
		v.GitHubToken = vdb.GitHubToken.String
	}
	if vdb.GitLabID.Valid {
		v.GitLabID = vdb.GitLabID.Int64
	}
	if vdb.GitLabToken.Valid {
		v.GitLabToken = vdb.GitLabToken.String
	}
	return nil
}

func (v *User) ToRow(vdb *db.User) error {
	vdb.ID = v.ID
	vdb.Name = v.Name
	vdb.SlackName = v.SlackName
	vdb.GitHubName.Valid = true
	vdb.GitHubName.String = v.GitHubName
	vdb.GitHubToken.Valid = true
	vdb.GitHubToken.String = v.GitHubToken
	vdb.GitLabID.Valid = true
	vdb.GitLabID.Int64 = int64(v.GitLabID)
	vdb.GitLabToken.Valid = true
	vdb.GitLabToken.String = v.GitLabToken
	return nil
}

//func (user User) setActionHistories() {
//
//	actions, err := db.LoadActions(user)
//
//	if err != nil {
//		log.Fatalf("Error load user actions: %q", err)
//		return
//	}
//
//	for _, action := range actions {
//		for _, actionHistory := range user.ActionHistories {
//			if actionHistory.isEqualDate(action.CreateAt) {
//				actionHistory.actions = append(actionHistory.actions, action)
//				break
//			}
//		}
//	}
//}

func (user User) Contributes() string {
	var turfs string

	for _, actionHistory := range user.ActionHistories {
		turfs = turfs + " " + actionHistory.turfColor()
	}

	return turfs
}

//// todayGitHubContributesCount is count today activity on GitHub
//func (user User) todayGitHubContributesCount() int {
//	var count int
//
//	for _, githubEvent := range user.GitHubEvents {
//		if githubEvent.isTodayContribute() && githubEvent.isCountable() {
//			count++
//		}
//	}
//
//	return count
//}
//
//// todayGitLabContributesCount is count today activity on GitLab
//func (user User) todayGitLabContributesCount() int {
//	var count int
//
//	for _, gitlabEvent := range user.GitLabEvents {
//		if gitlabEvent.isTodayContribute() {
//			count++
//		}
//	}
//
//	return count
//}
//
//// TodayContributesCount is count today activity on GitHub and GitLab
//func (user User) TodayContributesCount() int {
//	return user.todayGitHubContributesCount() + user.todayGitLabContributesCount()
//}
