package job

import "github.com/robfig/cron"

func BeginJob() {
	c := cron.New()

	//go c.AddFunc("*/10 * * * * ?", checkServeList)
	go c.AddFunc("0 0 1 * * ?", deleteOperateRecord)
	c.Start()
}
