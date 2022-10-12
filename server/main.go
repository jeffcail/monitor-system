package main

import (
	"net/http"

	"github.com/c/server-monitoring/server/internal/v1/job"
	"github.com/c/server-monitoring/server/router"

	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe(":9999", nil)
	}()

	job.BeginJob()
	router.RunServer()
}
