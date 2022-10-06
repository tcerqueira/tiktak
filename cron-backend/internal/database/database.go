package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	logger "github.com/tcerqueira/tiktak/cron-backend/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		PreferSimpleProtocol: false,
	}), &gorm.Config{})
	if err != nil {
		logger.Error.Fatalln("Could not connect to database")
	}

	GetConnection = func() *gorm.DB {
		return conn
	}
	logger.Info.Println("Connected to database")
}
