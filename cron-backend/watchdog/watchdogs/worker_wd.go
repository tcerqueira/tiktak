package watchdog

import (
	"time"

	database "github.com/tcerqueira/tiktak/cron-backend/internal/database"
	model "github.com/tcerqueira/tiktak/cron-backend/internal/model"
)

type WorkerWatchdog struct {
	WorkerTimeout   time.Duration
	PollingInterval time.Duration
}

func (ww *WorkerWatchdog) Start() {
	db := database.GetConnection()
	for {
		db.Delete(&model.CronWorker{}, "now() - updated_at > ?", ww.WorkerTimeout)
		time.Sleep(ww.PollingInterval)
	}
}
