package model

import "github.com/tcerqueira/tiktak/cron-backend/internal/database"

func init() {
	database.GetConnection().AutoMigrate(&CronJob{})
}

type CronJob struct {
	ID       string `gorm:"type:uuid;default:uuid_generate_v4()"`
	JobID    string
	Job      Job `gorm:"constraint:OnDelete:CASCADE;"`
	WorkerID string
	Worker   CronWorker `gorm:"constraint:OnDelete:CASCADE;"`
	Status   int        `gorm:"default:0"` // 0-new; 1-assigned
}
