package models

//CategoryFeed category_feeds
type CategoryFeed struct {
	CategoryID uint
	Category   Category
	FeedID     uint
	Feed       Feed
}
