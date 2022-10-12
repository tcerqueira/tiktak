package main

import (
	"sync"
	"time"

	watchdog "github.com/tcerqueira/tiktak/cron-backend/watchdog/watchdogs"
)

func main() {
	workerWd := watchdog.WorkerWatchdog{
		WorkerTimeout:   5 * time.Second,
		PollingInterval: time.Second,
	}
	jobWd := watchdog.JobWatchdog{
		PollingInterval: time.Second,
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go workerWd.Start()
	go jobWd.Start()

	wg.Wait()
}
