package main

import (
	cronsv "github.com/tcerqueira/tiktak/cron-backend/worker/cron-worker"
)

func main() {
	server := cronsv.NewServer()
	server.Start()
}
