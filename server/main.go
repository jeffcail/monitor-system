package main

import (
	"bz.service.cloud.monitoring/server/internal/v1/job"
	"bz.service.cloud.monitoring/server/router"
)

func main() {
	job.BeginJob()
	router.RunServer()
}
