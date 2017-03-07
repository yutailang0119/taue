package jobs

// ReportTaue is report activities which on git to slack
func ReportTaue() {

	targetUsers := loadUsersFromJSON()

	users := getContributes(targetUsers)

	postSlack(users)

}
