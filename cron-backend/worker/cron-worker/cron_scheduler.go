package cronsv

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tcerqueira/tiktak/cron-backend/internal/logger"
	"github.com/tcerqueira/tiktak/cron-backend/internal/model"
	"gopkg.in/robfig/cron.v2"
)

type CronScheduler struct {
	scheduler   *cron.Cron
	cronJobsMap map[string]cron.EntryID
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
	entryID := s.cronJobsMap[id]
	s.scheduler.Remove(entryID)
	delete(s.cronJobsMap, id)
}

func Trigger(j model.Job) {
	// logger.Info.Println("Triggered job: ", j.ID)
	client := http.DefaultClient
	var (
		urlStr     string
		bodyReader io.Reader
	)

	if j.WebhookMethod == "GET" {
		// Insert body param in URL in case of GET method
		urlStr = fmt.Sprintf(`%s?body=%s`, j.WebhookURL, url.QueryEscape(j.Body))
		bodyReader = nil
	} else {
		// Form body in every other method
		urlStr = j.WebhookURL
		jsonBody := fmt.Sprintf(`{"body":"%s"}`, j.Body) // JSON encoding "à pedreiro"
		bodyReader = bytes.NewReader([]byte(jsonBody))
	}

	req, err := http.NewRequest(j.WebhookMethod, urlStr, bodyReader)
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
