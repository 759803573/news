package models

import (
	"news/config"
	"time"
)

//ItemStatus ItemStatus
type ItemStatus struct {
	ItemID       uint
	Item         Item
	UserID       uint
	User         User
	CollectionID uint
	Collection   Collection
	ReadAt       time.Time
	LaterReadAt  time.Time
}

//Create Create
func (iss *ItemStatus) Create() {
	config.DB.Conn.Create(iss)
}
