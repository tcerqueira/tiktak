package cronwk

import (
	"sync"
	"time"

	database "github.com/tcerqueira/tiktak/cron-backend/internal/database"
	model "github.com/tcerqueira/tiktak/cron-backend/internal/model"
	cron "gopkg.in/robfig/cron.v2"
)

func init() {
	scheduler.Start()
}

var scheduler = cron.New()
var cronJobsMap = make(map[string]*model.CronJob)

func Init(cw *model.CronWorker) {
	db := database.GetConnection()

	// db.AutoMigrate(cw)
	db.Create(cw)
}

func Start(cw *model.CronWorker) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			cw.KeepAlive()
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Wait()
}

// func (cj *CronJob) Start() {
// 	if cj.running {
// 		return
// 	}
// 	job := cj.Job
// 	var err error

// 	exp := fmt.Sprintf("TZ=%s %s", job.Timezone, job.CronExpression)
// 	cj.EntryID, err = scheduler.AddFunc(exp, func() {
// 		job.Trigger()
// 	})
// 	if err != nil {
// 		logger.Error.Println("error 'Start': Adding Job: ", err.Error(), *job)
// 		return
// 	}

// 	cj.running = true
// 	cronJobsMap[job.ID] = cj
// 	// for key, el := range cronJobsMap {
// 	// 	fmt.Printf("%d => {%v, %d}\n", int(key), el.Job, int(el.EntryID))
// 	// }
// }

// func (j *Job) Trigger() {
// 	// fmt.Printf("%s - %d - %v\n", time.Now().Format(time.RubyDate), j.ID, j)
// 	client := http.DefaultClient
// 	var (
// 		url        string
// 		bodyReader io.Reader
// 	)

// 	if j.WebhookMethod == "GET" {
// 		// Insert body param in URL in case of GET method
// 		url = fmt.Sprintf(`%s?body="%s"`, j.WebhookURL, j.Body)
// 		bodyReader = nil
// 	} else {
// 		// Form body in every other method
// 		url = j.WebhookURL
// 		jsonBody := fmt.Sprintf(`{"body":"%s"}`, j.Body) // JSON encoding "à pedreiro"
// 		bodyReader = bytes.NewReader([]byte(jsonBody))
// 	}

// 	req, err := http.NewRequest(j.WebhookMethod, url, bodyReader)
// 	if err != nil {
// 		logger.Warn.Println("error 'Trigger': Creating request: ", err.Error(), *j)
// 		return
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	_, err = client.Do(req)
// 	if err != nil {
// 		logger.Warn.Println("error 'Trigger': Sending request: ", err.Error(), *j)
// 		return
// 	}

// 	return
// }
