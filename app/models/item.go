package models

import (
	"news/config"
	"time"
)

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
