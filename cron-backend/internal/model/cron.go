package model

import "github.com/tcerqueira/tiktak/cron-backend/internal/database"

func init() {
	database.GetConnection().AutoMigrate(&CronJob{})
}

type CronJob struct {
	JobID    string
	Job      Job `gorm:"primaryKey;constraint:OnDelete:CASCADE;"`
	WorkerID string
	Worker   CronWorker `gorm:"primaryKey;constraint:OnDelete:CASCADE;"`
	Status   int        `gorm:"default:0"` // 0-new; 1-assigned
}
