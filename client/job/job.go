package job

import "github.com/robfig/cron"

// BeginJob
func BeginJob() {
	c := cron.New()

	c.AddFunc("*/1200 * * * * ?", CheckClientVersion)
	c.AddFunc("*/30 * * * * ?", PushClientCpuPercent)
	c.AddFunc("*/30 * * * * ?", PushClientMemPercent)
	c.AddFunc("*/30 * * * * ?", PushClientDiskPercent)

	c.Start()
}
