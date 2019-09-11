package models

import (
	"news/config"

	"github.com/jinzhu/gorm"
)

const KeyFeedID = "feed_id"

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

//Items Items
func (f *Feed) Items(association *gorm.DB) *gorm.DB {
	if association == nil {
		association = config.DB.Conn.Debug().Model(f)
	}
	return association.
		Joins("left join items on feeds.id = items.feed_id").Where(f)
}

//GetItems Get items
func (f *Feed) GetItems(item *Item, association *gorm.DB) (items []*Item) {
	items = make([]*Item, 0)
	f.Items(association).Select("items.*").Scan(&items).Where(item)
	return
}
