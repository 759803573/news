package workers

import (
	"net/http"
	"time"
)

//FeedDownloadWork 负责下载
type FeedDownloadWork struct {
	FeedLink string
	client   http.Client
}

//Init 初始化
func (work *FeedDownloadWork) Init() *FeedDownloadWork {
	work.client = http.Client{Timeout: 10 * time.Second}
	return work
}

//Run 下载
func (work *FeedDownloadWork) Run() {
	if resp, err := work.client.Get(work.FeedLink); err == nil {
		(&FeedParseWorker{}).Init(work.FeedLink).Run(resp.Body)
	}
}
