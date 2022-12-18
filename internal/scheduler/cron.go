package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
)

var Scheduler *gocron.Scheduler

func Start() {
	Scheduler = gocron.NewScheduler(time.Local)
}
