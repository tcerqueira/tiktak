package cronsv

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/tcerqueira/tiktak/cron-backend/internal/logger"
	"github.com/tcerqueira/tiktak/cron-backend/internal/model"
	"gopkg.in/robfig/cron.v2"
)

type CronScheduler struct {
	scheduler   *cron.Cron
	cronJobsMap map[string]cron.EntryID
	mtx         sync.Mutex
}

func NewScheduler() *CronScheduler {
	var s CronScheduler
	s.scheduler = cron.New()
	s.cronJobsMap = make(map[string]cron.EntryID)

	return &s
}

func (s *CronScheduler) Start() {
	s.scheduler.Start()
}

func (s *CronScheduler) AddCronJob(job *model.Job) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	exp := fmt.Sprintf("TZ=%s %s", job.Timezone, job.CronExpression)
	entryID, err := s.scheduler.AddFunc(exp, func() {
		Trigger(*job)
	})
	if err != nil {
		logger.Error.Println("'Start': Adding Job: ", err.Error(), *job)
		return
	}
	s.cronJobsMap[job.ID] = entryID
}

func (s *CronScheduler) RemoveCronJob(id string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	entryID := s.cronJobsMap[id]
	s.scheduler.Remove(entryID)
	delete(s.cronJobsMap, id)
}

func Trigger(j model.Job) {
	// fmt.Printf("%s - %d - %v\n", time.Now().Format(time.RubyDate), j.ID, j)
	client := http.DefaultClient
	var (
		url        string
		bodyReader io.Reader
	)

	if j.WebhookMethod == "GET" {
		// Insert body param in URL in case of GET method
		url = fmt.Sprintf(`%s?body="%s"`, j.WebhookURL, j.Body)
		bodyReader = nil
	} else {
		// Form body in every other method
		url = j.WebhookURL
		jsonBody := fmt.Sprintf(`{"body":"%s"}`, j.Body) // JSON encoding "à pedreiro"
		bodyReader = bytes.NewReader([]byte(jsonBody))
	}

	req, err := http.NewRequest(j.WebhookMethod, url, bodyReader)
	if err != nil {
		logger.Warn.Println("'Trigger': Creating request: ", err.Error(), j)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	_, err = client.Do(req)
	if err != nil {
		logger.Warn.Println("'Trigger': Sending request: ", err.Error(), j)
		return
	}

	return
}
