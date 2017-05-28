package jobs

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/Yu-taro/taue/taue/models"
)

const (
	DB_USER     = ""
	DB_PASSWORD = ""
	DB_NAME     = ""
	DB_HOST     = ""
	DB_PORT     = ""
)

// loadUsers is that load users profile from DataBase
func loadUsers() (users []models.User, err error) {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME, DB_HOST, DB_PORT)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE del_flag = FALSE ")
	if err != nil {
		log.Fatalf("Error query : %q", err)
		return nil, err
	}

	for rows.Next() {
		var user models.User
		var delFlag bool
		err = rows.Scan(&user.ID, &user.Name, &user.SlackName, &user.GitHubName, &user.GitHubToken, &user.GitLabID, &user.GitLabToken, &delFlag)
		if err != nil {
			log.Fatalf("Error scan : %q", err)
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil

}
