package models

import (
	"news/config"
	"time"
)

const KeyItemID = "item_id"

//Item Item
type Item struct {
	config.DBBaseModel
	FeedID      uint
	Feed        Feed
	Title       string
	Description string
	Content     string `gorm:"Type:text"`
	Link        string `gorm:"UNIQUE_INDEX;NOT NULL"`
	PublishAt   time.Time
	Author      string
	GUID        string `gorm:"UNIQUE_INDEX"`
	Image       string
	Enclosures  string
	Categories  string
}

//Create Create
func (item *Item) Create() {
	config.DB.Conn.Create(item)
}

//CreateOrIgnore CreateOrIgnore
func (item *Item) CreateOrIgnore() {
	tmpItem := &Item{}

	if config.DB.Conn.Debug().Where("link = ?", item.Link).First(tmpItem).RecordNotFound() {
		item.Create()
	} else {
		*item = *tmpItem
	}
}
