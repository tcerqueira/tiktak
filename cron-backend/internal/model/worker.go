package model

import (
	"time"

	"github.com/tcerqueira/tiktak/cron-backend/internal/database"
)

func init() {
	database.GetConnection().AutoMigrate(&CronWorker{})
}

type CronWorker struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4()"`
	AlivePing bool
	UpdatedAt time.Time
	WorkCount int
	Ready     bool
}

func (cw *CronWorker) KeepAlive() {
	db := database.GetConnection()
	cw.AlivePing = !cw.AlivePing
	db.Model(cw).Update("alive_ping", cw.AlivePing)
}
