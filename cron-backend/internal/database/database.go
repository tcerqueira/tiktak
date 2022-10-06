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

var (
	conn          *gorm.DB
	GetConnection func() *gorm.DB
)

func init() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  GetDSN(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		logger.Error.Fatalln("Could not connect to database")
	}

	GetConnection = func() *gorm.DB {
		return conn
	}
	logger.Info.Println("Connected to database")
}

func GetDSN() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("SUPABASE_DB_USER"),
		os.Getenv("SUPABASE_DB_PASS"),
		os.Getenv("SUPABASE_DB_HOST"),
		os.Getenv("SUPABASE_DB_POOL_PORT"),
		os.Getenv("SUPABASE_DB_DBNAME"),
	)
}

func GetURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("SUPABASE_DB_USER"),
		os.Getenv("SUPABASE_DB_PASS"),
		os.Getenv("SUPABASE_DB_HOST"),
		os.Getenv("SUPABASE_DB_POOL_PORT"),
		os.Getenv("SUPABASE_DB_DBNAME"),
	)
}

func GetDirectDSN() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("SUPABASE_DB_USER"),
		os.Getenv("SUPABASE_DB_PASS"),
		os.Getenv("SUPABASE_DB_HOST"),
		os.Getenv("SUPABASE_DB_PORT"),
		os.Getenv("SUPABASE_DB_DBNAME"),
	)
}

func GetDirectURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("SUPABASE_DB_USER"),
		os.Getenv("SUPABASE_DB_PASS"),
		os.Getenv("SUPABASE_DB_HOST"),
		os.Getenv("SUPABASE_DB_PORT"),
		os.Getenv("SUPABASE_DB_DBNAME"),
	)
}
