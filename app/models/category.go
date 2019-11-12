package models

import (
	"news/app/helpers"
	"news/config"

	"github.com/jinzhu/gorm"
)

const KeyCategoryID = "category_id"

//Category feed category
type Category struct {
	config.DBBaseModel
	Name   string
	UserID uint
	User   User
}

type CategoryStatus struct {
	Name        string `json:"category_name"`
	UnReadCount int    `json:"unread_count"`
}

//First First
func (c *Category) First() (err error) {
	tmpCategory := &Category{}
	if !config.DB.Conn.Debug().Where("user_id = ?", c.UserID).First(tmpCategory).RecordNotFound() {
		*c = *tmpCategory
	} else {
		err = helpers.ErrRecordNotFound
	}
	return
}

//FindAll FindAll
func (c *Category) FindAll() (categories []Category, err error) {
	if config.DB.Conn.Debug().Where("user_id = ?", c.UserID).Find(&categories).RecordNotFound() {
		err = helpers.ErrRecordNotFound
	}
	return
}

//Feeds 定义关联 Feeds 的方式
func (c *Category) Feeds(association *gorm.DB) *gorm.DB {
	if association == nil {
		association = config.DB.Conn.Debug().Model(c)
	}
	return association.
		Joins("left join category_feeds  on category_feeds.category_id = categories.id").
		Joins("left join feeds on feeds.id = category_feeds.feed_id").Where(c)
}

//GetFeeds Get Feeds
func (c *Category) GetFeeds(feed *Feed, association *gorm.DB) (feeds []*Feed) {
	feeds = make([]*Feed, 0)
	c.Feeds(association).Select("feeds.*").Where(feed).Scan(&feeds)
	return
}

//GetItems Get Items
func (c *Category) GetItems(feed *Feed, association *gorm.DB) (items []*Item) {
	items = (&Feed{}).GetItems(&Item{}, c.Feeds(association))
	return
}

//GetStatus GetStatus
func (c *Category) GetStatus(association *gorm.DB) (categoryStatus []CategoryStatus) {
	if true {
		categoryStatus = make([]CategoryStatus, 0)
		(&Item{}).ItemStatus((&Feed{}).Items(c.Feeds(association))).
			Select("categories.name, sum(item_statuses.read_at is null) as un_read_count").
			Group("categories.name").Scan(&categoryStatus)
	} else {
		_ = helpers.ErrRecordNotFound
	}
	return
}
