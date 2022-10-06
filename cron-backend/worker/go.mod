module github.com/tcerqueira/tiktak/cron-backend/worker

go 1.19

replace github.com/tcerqueira/tiktak/cron-backend/internal => ../internal

require (
	github.com/lib/pq v1.10.2
	github.com/tcerqueira/tiktak/cron-backend/internal v0.0.0-20221006133016-9b5a1b20b668
	gopkg.in/robfig/cron.v2 v2.0.0-20150107220207-be2e0b0deed5
)

require (
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/text v0.3.7 // indirect
	gorm.io/driver/postgres v1.3.10 // indirect
	gorm.io/gorm v1.23.10 // indirect
)
