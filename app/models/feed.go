package models

import "news/config"

//Feed feed
type Feed struct {
	config.DBBaseModel
	Name         string
	UserID       uint
	User         User
	Title        string
	Description  string
	FeedLink     string
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
