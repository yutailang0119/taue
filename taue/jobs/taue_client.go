package jobs

// Report taue is report activities which on git to slack
func ReportTaue() {

	targetUsers, err := loadUsers()

	if err != nil {
		return
	}

	users := getContributes(targetUsers)

	postSlack(users)

}
