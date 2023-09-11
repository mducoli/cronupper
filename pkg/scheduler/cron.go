package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/mducoli/cronupper/pkg/executer"
	"github.com/mducoli/cronupper/pkg/types"
)

func CronJobBlockingLog(jobs map[string]types.Job) {
	s := gocron.NewScheduler(time.UTC)

	// fmt.Printf("jobs: %+v\n", jobs)

	for k := range jobs {

		log.Printf(`Scheduling job: "%v"`, k)

		_, err := s.Cron(jobs[k].Cron).Do(executeWrapper(jobs[k]))

		if err != nil {
			log.Println(err)
		}
	}

	log.Println("All jobs scheduled")

	s.StartBlocking()
}

func executeWrapper(job types.Job) func() {
	return func() {
		log.Printf(`Executing job: "%v"`, job.Id)

		err := executer.Execute(&job)
		if err != nil {
			log.Println(err)
		}
	}
}
