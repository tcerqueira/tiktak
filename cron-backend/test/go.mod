module github.com/tcerqueira/tiktak/cron-backend/test

go 1.19

replace github.com/tcerqueira/tiktak/cron-backend/internal => ../internal

require (
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/tcerqueira/tiktak/cron-backend/internal v0.0.0-00010101000000-000000000000
)

require github.com/felixge/httpsnoop v1.0.1 // indirect
