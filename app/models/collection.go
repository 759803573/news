package models

import (
	"news/config"
)

//Collection Collection
type Collection struct {
	config.DBBaseModel
	UserID uint
	User   User
	Name   string
	Count  int
}
