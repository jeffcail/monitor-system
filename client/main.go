package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/c/monitor-system/client/job"

	"github.com/c/monitor-system/client/bootstarp"
	"github.com/c/monitor-system/client/router"

	"github.com/kardianos/service"
)

type program struct {
	log service.Logger
	cfg *service.Config
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {

	bootstarp.InitBoot()
	job.BeginJob()
	router.RunClientServer()

	wg.Done()
}

func (p *program) Stop(s service.Service) error {
	return nil
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	svcConfig := &service.Config{
		Name:        "client-monitor",
		DisplayName: "client-monitor",
		Description: "client-monitor",
	}
	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			x := s.Install()
			if x != nil {
				fmt.Println("error:", x.Error())
				return
			}
			fmt.Println("服务安装成功")
			return
		} else if os.Args[1] == "uninstall" {
			x := s.Uninstall()
			if x != nil {
				fmt.Println("error:", x.Error())
				return
			}
			fmt.Println("服务卸载成功")
			return
		}
	}
	err = s.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	wg.Wait()
}
