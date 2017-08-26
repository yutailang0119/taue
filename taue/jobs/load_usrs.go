package jobs

import (
	"log"
	_ "github.com/lib/pq"

	"github.com/yutailang0119/taue/taue/models"
	"github.com/yutailang0119/taue/taue/db"
)

// loadUsers is that load users profile from DataBase
func loadUsers() (users []models.User, err error) {

	database, err := db.OpenDB()

	if err != nil {
		log.Fatalf("Error opening database: %q", err)
		return nil, err
	}

	defer database.Close()

	userList := db.UserList{}

	err = userList.Load(database)

	if err != nil {
		log.Fatalf("Error load users: %q", err)
		return nil, err
	}

	return users, nil

}
