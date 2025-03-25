package main

import (
	"fmt"
	"github.com/h-varmazyar/p3o/migrations"
	log "github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
)

func main() {
	dep, err := InjectDependencies()
	if err != nil {
		log.Panic(err)
	}

	signal.Notify(dep.StopSignal, syscall.SIGINT, syscall.SIGTERM)

	if err = migrations.Migrate(dep.DB); err != nil {
		log.Panic(err)
	}

	address := fmt.Sprintf("%v:%v", dep.Cfg.Server.HttpAddress, dep.Cfg.Server.HttpPort)
	if err := dep.Routers.Router.StartServing(dep.Gin, address); err != nil {
		log.Panic(err)
	}

	<-dep.StopSignal
}
