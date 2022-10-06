package main

import (
	logger "github.com/tcerqueira/tiktak/cron-backend/internal/logger"
	"github.com/tcerqueira/tiktak/cron-backend/internal/model"
	cronwk "github.com/tcerqueira/tiktak/cron-backend/worker/cron-worker"
)

func main() {
	worker := model.CronWorker{}
	cronwk.Init(&worker)

	logger.Info.Println("Starting worker ", worker.ID)
	cronwk.Start(&worker)
}
