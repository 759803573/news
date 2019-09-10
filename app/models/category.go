package models

import (
	"news/app/helpers"
	"news/config"
)

const KeyCategoryID = "category_id"

//Category feed category
type Category struct {
	config.DBBaseModel
	Name   string
	UserID uint
	User   User
	Feeds  []Feed `gorm:"many2many:category_feeds;"`
}

type CategoryStatus struct {
	Name        string `json:"category_name"`
	UnReadCount int    `json:"unread_count"`
}

//GetByUserID GetByUserID
func (c *Category) GetByUserID() (err error) {
	tmpCategory := &Category{}
	if !config.DB.Conn.Debug().Where("user_id = ?", c.UserID).First(tmpCategory).RecordNotFound() {
		*c = *tmpCategory
	} else {
		err = helpers.ErrRecordNotFound
	}
	return
}

//GetStatus GetStatus
func (c *Category) GetStatus() []CategoryStatus {
	// tmpCategory := &Category{}
	if true {
		categoryStatus := make([]CategoryStatus, 0)
		config.DB.Conn.Debug().Model(c).
			Select("categories.name, sum(iss.read_at is null) as un_read_count").
			Joins("left join category_feeds cf on cf.category_id = categories.id").
			Joins("left join feeds f on f.id = cf.feed_id").
			Joins("left join items i on i.feed_id = f.id").
			Joins("left join item_statuses iss on iss.item_id = i.id").
			Group("categories.name").Scan(&categoryStatus)
		return categoryStatus
	} else {
		_ = helpers.ErrRecordNotFound
	}
	return nil
}
