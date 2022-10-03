package cronjob

import (
	"fmt"

	cron "gopkg.in/robfig/cron.v2"
)

var scheduler = cron.New()
var cronJobsMap = make(map[JobID]*CronJob)

type CronJobID cron.EntryID

type CronJob struct {
	Job     *Job
	EntryID cron.EntryID
	started bool
}

func init() {
	scheduler.Start()
}

func (cj *CronJob) Start() {
	if cj.started {
		return
	}
	job := cj.Job
	var err error

	exp := fmt.Sprintf("TZ=%s %s", job.Timezone, job.CronExpression)
	cj.EntryID, err = scheduler.AddFunc(exp, func() {
		job.Trigger()
	})
	if err != nil {
		fmt.Println("error 'Start': Adding Job: ", err.Error(), *job)
		return
	}

	cj.started = true
	cronJobsMap[job.ID] = cj
	// for key, el := range cronJobsMap {
	// 	fmt.Printf("%d => {%v, %d}\n", int(key), el.Job, int(el.EntryID))
	// }
}
