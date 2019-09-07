package workers

var fedsWork *FeedsSyncWork

//GetFeedsSyncWork 管理所有同步任务
func GetFeedsSyncWork() *FeedsSyncWork {
	if fedsWork == nil {
		fedsWork = &FeedsSyncWork{status: ""}
		fedsWork.Init()
	}

	return fedsWork
}

//FeedsSyncWork feedsSyncWork
type FeedsSyncWork struct {
	status   string
	feedList []string
}

//Init init
func (work *FeedsSyncWork) Init() {
	if work.status == "" {
		work.updateFeedList()
		work.status = "start"
	}
}

func (work *FeedsSyncWork) updateFeedList() {
	work.feedList = []string{
		"http://www.sspai.com/feed",
		"http://www.zhihu.com/rss",
		"http://www.quora.com/rss",
	}
}

//Run run
func (work *FeedsSyncWork) Run() {
	for _, feed := range work.feedList {
		(&FeedDownloadWork{FeedLink: feed}).Init().Run()
	}
}
