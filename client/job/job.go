package job

import "github.com/robfig/cron"

// BeginJob
func BeginJob() {
	c := cron.New()

	c.AddFunc("*/1200 * * * * ?", CheckClientVersion)

	c.Start()
}
