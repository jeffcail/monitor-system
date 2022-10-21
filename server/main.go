package main

import (
	"net/http"

	"github.com/c/monitor-system/server/internal/v1/job"
	"github.com/c/monitor-system/server/router"

	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe(":9999", nil)
	}()

	job.BeginJob()
	router.RunServer()
}
