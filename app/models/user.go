package models

import "news/config"

//User user model
type User struct {
	config.DBBaseModel
	Token      string
	Categories []Category
}
