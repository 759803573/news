package models

//Category feed category
type Category struct {
	Name   string
	UserID uint
	User   User
	FeedID uint
	Feed   Feed
}
