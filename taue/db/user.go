package db

import (
	"database/sql"
	"log"
)

// User define a user from Database
type User struct {
	ID          int
	Name        string
	SlackName   string
	GitHubName  sql.NullString
	GitHubToken sql.NullString
	GitLabID    sql.NullInt64
	GitLabToken sql.NullString
}

type UserList []User

// Load is that load users profile from DataBase
func (v *UserList) Load(db *sql.DB) error {

	rows, err := db.Query("SELECT * FROM users WHERE del_flag = FALSE")

	if err != nil {
		log.Fatalf("Error users query : %q", err)
		return err
	}

	if err := v.FromRows(rows); err != nil {
		return err
	}

	return nil

}

func (v *UserList) FromRows(rows *sql.Rows) error {
	var users []User
	for rows.Next() {
		vdb := User{}
		var delFlag bool

		err := rows.Scan(&vdb.ID, &vdb.Name, &vdb.SlackName, &vdb.GitHubName, &vdb.GitHubToken, &vdb.GitLabID, &vdb.GitLabToken, &delFlag)

		if err != nil {
			log.Fatalf("Error user scan : %q", err)
			return err
		}

		users = append(users, vdb)

	}

	*v = users
	return nil
}
