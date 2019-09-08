package models

import (
	"news/app/helpers"
	"news/config"
)

//Category feed category
type Category struct {
	Name   string
	UserID uint
	User   User
	FeedID uint
	Feed   Feed
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
