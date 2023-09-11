package main

import (
	"github.com/mducoli/cronupper/pkg/config"
	"github.com/mducoli/cronupper/pkg/scheduler"
)

func main() {

	config, err := config.Parse()
	if err != nil {
		panic(err)
	}

  scheduler.CronJobBlockingLog(config.Jobs)
}
