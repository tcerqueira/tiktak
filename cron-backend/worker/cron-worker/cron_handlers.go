package cronsv

import (
	"encoding/json"

	"github.com/lib/pq"
	"github.com/tcerqueira/tiktak/cron-backend/internal/logger"
	"github.com/tcerqueira/tiktak/cron-backend/internal/model"
)

func HandleCreateCron(cs *CronServer, event *pq.Notification) {
	logger.Info.Println("Event 'Create': ", event.Extra)
	var newJob model.Job
	err := json.Unmarshal([]byte(event.Extra), &newJob)
	if err != nil {
		logger.Error.Println("'Handle create': ", err.Error(), event.Extra)
		return
	}
	cs.AddCronJob(&newJob)
}

func HandleDeleteCron(cs *CronServer, event *pq.Notification) {
	logger.Info.Println("Event 'Delete': ", event.Extra)
	cs.RemoveCronJob(event.Extra)
}
