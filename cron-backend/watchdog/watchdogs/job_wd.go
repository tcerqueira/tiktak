package watchdog

import (
	"time"

	database "github.com/tcerqueira/tiktak/cron-backend/internal/database"
	"github.com/tcerqueira/tiktak/cron-backend/internal/model"
	"gorm.io/gorm"
)

type JobWatchdog struct {
	PollingInterval time.Duration
}

func (ww *JobWatchdog) Start() {
	db := database.GetConnection()
	for {
		db.Transaction(func(tx *gorm.DB) error {
			var orfanJobs []model.Job
			db.Model(&model.Job{}).
				Select("jobs.*").
				Joins("LEFT JOIN cron_jobs ON cron_jobs.job_id=jobs.id").
				Where("cron_jobs.worker_id ISNULL").
				Scan(&orfanJobs)

			if len(orfanJobs) > 0 {
				result := db.Exec("DELETE FROM jobs WHERE jobs.id IN (SELECT jobs.id FROM jobs LEFT JOIN cron_jobs ON cron_jobs.job_id=jobs.id WHERE cron_jobs.worker_id ISNULL)")
				if result.Error != nil {
					return result.Error
				}
				if err := db.Create(orfanJobs).Error; err != nil {
					return err
				}
			}

			return nil
		})
		time.Sleep(ww.PollingInterval)
	}
}
