package cronjob

import "voz/global"

func RunCronjob() {
	//CrawlThreads(global.F17, "thread")
	CrawlComments(global.TestThread, "comment")
}
