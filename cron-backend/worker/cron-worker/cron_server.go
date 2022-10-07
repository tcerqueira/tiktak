package cronsv

import (
	"sync"
	"time"

	pq "github.com/lib/pq"
	database "github.com/tcerqueira/tiktak/cron-backend/internal/database"
	"github.com/tcerqueira/tiktak/cron-backend/internal/logger"
	model "github.com/tcerqueira/tiktak/cron-backend/internal/model"
)

type CronServer struct {
	CronWorker     model.CronWorker
	Scheduler      *CronScheduler
	createListener *pq.Listener
	deleteListener *pq.Listener
}

func NewServer() *CronServer {
	var cs CronServer
	cs.CronWorker = model.CronWorker{}
	cs.Scheduler = NewScheduler()

	result := database.GetConnection().Create(&cs.CronWorker)
	if result.Error != nil {
		panic(result.Error.Error())
	}

	return &cs
}

func (cs *CronServer) Start() {
	var err error
	cs.createListener, err = registerListener("create_" + cs.CronWorker.ID)
	if err != nil {
		panic(err.Error())
	}
	cs.deleteListener, err = registerListener("delete_" + cs.CronWorker.ID)
	if err != nil {
		panic(err.Error())
	}
	cs.Scheduler.Start()

	var wg sync.WaitGroup
	wg.Add(3)
	go cs.listenChannel(cs.createListener.Notify, HandleCreateCron)
	go cs.listenChannel(cs.deleteListener.Notify, HandleDeleteCron)
	go cs.heartBeat()
	cs.Ready()

	logger.Info.Println("Starting cron server ", cs.CronWorker.ID)
	wg.Wait()
}

func (cs *CronServer) heartBeat() {
	for {
		cs.CronWorker.KeepAlive()
		time.Sleep(1 * time.Second)
	}
}

func (cs *CronServer) listenChannel(channel chan *pq.Notification, handler func(*CronServer, *pq.Notification)) {
	for {
		select {
		case event := <-channel:
			go handler(cs, event)
		}
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
	listener := pq.NewListener(database.GetDirectDSN(), minReconn, maxReconn, report)
	err := listener.Listen(channel)

	return listener, err
}

func (cs *CronServer) Ready() {
	cs.CronWorker.Ready = true
	database.GetConnection().Save(&cs.CronWorker)
}
