package model

import (
	database "github.com/tcerqueira/tiktak/cron-backend/internal/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func init() {
	database.GetConnection().AutoMigrate(&Job{})
}

type Job struct {
	ID             string `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	WebhookURL     string `json:"webhook_url"`
	WebhookMethod  string `json:"webhook_method"`
	Body           string `json:"body"`
	CronExpression string `json:"cron_expression"`
	Timezone       string `json:"timezone"`
}

func FetchAllJobs(rows *[]Job) *gorm.DB {
	db := database.GetConnection()
	return db.Find(&rows)
}

func FetchJob(job *Job) *gorm.DB {
	db := database.GetConnection()
	result := db.Find(job)
	if result.RowsAffected == 0 {
		job.ID = ""
	}
	return result
}

func InsertJob(job *Job) *gorm.DB {
	db := database.GetConnection()
	return db.Create(&job)
}

func UpdateJob(target, job *Job) *gorm.DB {
	db := database.GetConnection()
	result := db.Model(target).Clauses(clause.Returning{}).Updates(job)
	if result.RowsAffected == 0 {
		target.ID = ""
	}
	return result
}

func DeleteJob(id string) *gorm.DB {
	db := database.GetConnection()
	return db.Delete(&Job{ID: id})
}
