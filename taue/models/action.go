package models

import (
	"time"

	"github.com/yutailang0119/taue/taue/db"
)

type Action struct {
	ID         int
	UserID     int
	SourceType string
	Count      int
	CreateAt   time.Time
}

func (v *Action) FromRow(vdb *db.Action) error {
	v.ID = vdb.ID
	v.UserID = vdb.UserID
	v.SourceType = vdb.SourceType
	v.Count = vdb.Count
	v.CreateAt = vdb.CreateAt
	return nil
}

func (v *Action) ToRow(vdb *db.Action) error {
	vdb.ID = v.ID
	vdb.UserID = v.UserID
	vdb.SourceType = v.SourceType
	vdb.Count = v.Count
	vdb.CreateAt = v.CreateAt
	return nil
}

type ActionList []Action
