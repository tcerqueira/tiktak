package gateways

import (
	"sync"
	"time"

	job "github.com/tcerqueira/tiktak/cron-backend/Dispatcher/pkg/Jobs"
)

type JobContainer struct {
	job       *job.Job
	createdAt time.Time
	mtx       sync.Mutex
}

var (
	// Serving as DB
	jobsMap map[job.JobID]JobContainer
)

func FetchJob(id job.JobID) (job.Job, error) {
	return job.Job{}, nil
}

func FetchJobsList() []job.Job {
	return nil
}

func InsertJob(j *job.Job) error {
	return nil
}

func UpdateJob(j *job.Job) error {
	return nil
}

func DeleteJob(id job.JobID) error {
	return nil
}
