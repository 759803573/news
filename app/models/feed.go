package models

import (
	"news/config"
)

//Feed feed
type Feed struct {
	config.DBBaseModel
	Name         string
	Title        string
	Description  string
	FeedLink     string `gorm:"unique;not null"`
	Link         string
	Author       string
	Language     string
	Image        string
	Copyright    string
	Generator    string
	CategoriesID uint
	Category     *Category
	Items        []Item
}

func (feed *Feed) Create() {
	config.DB.Conn.Create(feed)
}

//CreateOrUpdate CreateOrUpdate
func (feed *Feed) CreateOrUpdate() {
	tmpFeed := &Feed{}

	if config.DB.Conn.Debug().Where("feed_link = ?", feed.FeedLink).First(tmpFeed).RecordNotFound() {
		feed.Create()
	} else {
		config.DB.Conn.Model(tmpFeed).UpdateColumns(feed)

		*feed = *tmpFeed
	}
}

//CreateItems CreateItems
func (feed *Feed) CreateItems(items []*Item) {
	tx := config.DB.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, item := range items {
		item.FeedID = feed.ID
		item.CreateOrIgnore()
	}
	tx.Commit()
}
