package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	logger "github.com/tcerqueira/tiktak/cron-backend/internal/logger"
	model "github.com/tcerqueira/tiktak/cron-backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var conn *gorm.DB
var GetConnection func() *gorm.DB

func init() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal(err)
	}
	pq_host := os.Getenv("SUPABASE_DB_HOST")
	pq_user := os.Getenv("SUPABASE_DB_USER")
	pq_password := os.Getenv("SUPABASE_DB_PASS")
	pq_dbname := os.Getenv("SUPABASE_DB_DBNAME")
	pq_port := os.Getenv("SUPABASE_DB_PORT")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		pq_user, pq_password, pq_host, pq_port, pq_dbname)
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		logger.Error.Fatalln("Could not connect to database")
	}

	conn.AutoMigrate(&model.Job{})
	GetConnection = func() *gorm.DB {
		return conn
	}
	logger.Info.Println("Connected to database")
}

func FetchAllJobs(rows *[]model.Job) *gorm.DB {
	db := GetConnection()
	return db.Find(&rows)
}

func FetchJob(job *model.Job) *gorm.DB {
	db := GetConnection()
	result := db.Find(job)
	if result.RowsAffected == 0 {
		job.ID = ""
	}
	return result
}

func InsertJob(job *model.Job) *gorm.DB {
	db := GetConnection()
	return db.Create(&job)
}

func UpdateJob(target, job *model.Job) *gorm.DB {
	db := GetConnection()
	result := db.Model(target).Clauses(clause.Returning{}).Updates(job)
	if result.RowsAffected == 0 {
		target.ID = ""
	}
	return result
}

func DeleteJob(id model.JobID) *gorm.DB {
	db := GetConnection()
	return db.Delete(&model.Job{}, id)
}
