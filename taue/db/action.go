package db

import (
	"time"
	"database/sql"
	"log"
)

type Action struct {
	ID         int
	UserID     int
	SourceType string
	Count      int
	CreateAt   time.Time
}

type ActionList []Action

func (v *ActionList) Load(db *sql.DB, userId int) error {

	rows, err := db.Query("SELECT * FROM actions WHERE user_id = $1", userId)

	if err != nil {
		log.Fatalf("Error actions query : %q", err)
	}

	if err := v.FromRows(rows); err != nil {
		return err
	}

	return nil

}

func (v *ActionList) FromRows(rows *sql.Rows) error {
	var actions []Action
	for rows.Next() {
		action := Action{}
		err := rows.Scan(&action.ID, &action.UserID, &action.SourceType, &action.Count, &action.CreateAt)
		if err != nil {
			log.Fatalf("Error action scan : %q", err)
		}

		actions = append(actions, action)
	}

	*v = actions
	return nil
}

func (v Action) Update(db *sql.DB) error {

	res, err := db.Exec("INSERT INTO actions(user_id, source_type, count) VALUES ($1, $2, $3)", v.UserID, v.SourceType, v.Count)

	if err != nil {
		log.Fatalf("Error save action: %q", err)
		return err
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Fatalf("Error None last insert id: %q", err)
		return err
	}
	print(id)
	return nil
}
