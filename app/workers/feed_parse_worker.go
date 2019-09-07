package workers

import (
	"fmt"
	"io"

	"news/app/models"

	"github.com/mmcdole/gofeed"
)

//FeedParseWorker 解析入库
type FeedParseWorker struct {
	fp       *gofeed.Parser
	feedLink string
}

//Init 初始化
func (work *FeedParseWorker) Init(feedLink string) *FeedParseWorker {
	work.fp = gofeed.NewParser()
	work.feedLink = feedLink
	return work
}

//Run 解析入库
func (work *FeedParseWorker) Run(feedData io.Reader) {
	fmt.Println(work.feedLink)
	feed, err := work.fp.Parse(feedData)
	if err != nil {
		fmt.Println(err)
		return
	}
	feedModel := models.Feed{
		FeedLink:    work.feedLink,
		Title:       feed.Title,
		Description: feed.Description,
		Link:        feed.Link,
		Language:    feed.Language,
		// Image:       feed.Image,
		Copyright: feed.Copyright,
	}
	(&feedModel).CreateOrUpdate()

	var items = make([]*models.Item, len(feed.Items))

	for idx, item := range feed.Items {
		items[idx] = &models.Item{
			Link:        item.Link,
			Title:       item.Title,
			Description: item.Description,
			Content:     item.Content,
			GUID:        item.GUID,
		}
	}
	feedModel.CreateItems(items)
}
