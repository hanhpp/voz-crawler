package cronjob

import (
	"github.com/fatih/color"
	"github.com/robfig/cron/v3"
	"time"
	"voz/config"
)

func CrawlComments(url string, fileName string) {
	color.Green("cron job: Crawling thread from %s\nSaving into file /text/%s.txt", url, fileName)
	skipLogger := config.SkipLogger{}
	c := cron.New(
		cron.WithLocation(time.UTC),
		cron.WithChain(cron.SkipIfStillRunning(skipLogger)),
	)
	_, _ = c.AddFunc("@every 0h0m10s", func() { // 23h59m GMT +8
		config.GetLogger().Info("Running crawler...")
		VisitAndCollectFromURL(url, fileName)
	})
	c.Start()
}
