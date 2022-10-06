package cronsv

import (
	"database/sql"
	"sync"
	"time"

	pq "github.com/lib/pq"
	database "github.com/tcerqueira/tiktak/cron-backend/internal/database"
	"github.com/tcerqueira/tiktak/cron-backend/internal/logger"
	model "github.com/tcerqueira/tiktak/cron-backend/internal/model"
	cron "gopkg.in/robfig/cron.v2"
)

type CronServer struct {
	cronWorker     model.CronWorker
	scheduler      *cron.Cron
	cronJobsMap    map[string]*model.CronJob
	createListener *pq.Listener
	deleteListener *pq.Listener
}

func NewServer() *CronServer {
	var cs CronServer
	cs.cronWorker = model.CronWorker{}
	cs.scheduler = cron.New()
	cs.cronJobsMap = make(map[string]*model.CronJob)

	result := database.GetConnection().Create(&cs.cronWorker)
	if result.Error != nil {
		panic(result.Error.Error())
	}

	_, err := sql.Open("postgres", database.GetDSN())
	if err != nil {
		panic(err)
	}

	return &cs
}

func (cs *CronServer) Start() {
	var err error
	cs.createListener, err = registerListener("create_" + cs.cronWorker.ID)
	if err != nil {
		panic(err.Error())
	}
	cs.deleteListener, err = registerListener("delete_" + cs.cronWorker.ID)
	if err != nil {
		panic(err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(3)
	go cs.HeartBeat()
	go listenChannel(cs.createListener.Notify)
	go listenChannel(cs.deleteListener.Notify)
	cs.Ready()

	logger.Info.Println("Starting cron server ", cs.cronWorker.ID)
	wg.Wait()
}

func (cs *CronServer) HeartBeat() {
	for {
		cs.cronWorker.KeepAlive()
		time.Sleep(1 * time.Second)
	}
}

func listenChannel(channel chan *pq.Notification) {
	for event := range channel {
		logger.Info.Println("Create event: ", event)
	}
}

func registerListener(channel string) (*pq.Listener, error) {
	report := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			logger.Info.Println(err.Error())
		}
	}
	minReconn := 10 * time.Second
	maxReconn := time.Minute
	listener := pq.NewListener(database.GetDSN(), minReconn, maxReconn, report)
	err := listener.Listen(channel)

	return listener, err
}

func (cs *CronServer) Ready() {
	cs.cronWorker.Ready = true
	database.GetConnection().Save(&cs.cronWorker)
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
