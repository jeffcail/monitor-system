package job

import "github.com/robfig/cron"

// BeginJob
func BeginJob() {
	c := cron.New()

	//c.AddFunc("*/1200 * * * * ?", CheckClientVersion)
	//c.AddFunc("*/2 * * * * ?", PushClientCpuPercent)
	//c.AddFunc("*/2 * * * * ?", PushClientMemPercent)
	//c.AddFunc("*/2 * * * * ?", PushClientDiskPercent)

	c.Start()
}
