package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	logger "github.com/tcerqueira/tiktak/cron-backend/internal/logger"
	model "github.com/tcerqueira/tiktak/cron-backend/internal/model"
)

type ResponsePayload struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func HandleGetJobsList(res http.ResponseWriter, req *http.Request) {
	logger.Info.Println("Request - jobs list")
	jobs := []model.Job{}
	result := model.FetchAllJobs(&jobs)

	writeResponse(res, jobs, result.Error)
}

func HandleGetJob(res http.ResponseWriter, req *http.Request) {
	id, err := uuid.Parse(mux.Vars(req)["id"])
	if ok := handleParamsError(res, err); !ok {
		return
	}
	idStr := id.String()
	logger.Info.Printf("Request - jobs by id (%s)\n", idStr)

	job := model.Job{ID: idStr}
	result := model.FetchJob(&job)
	if job.ID == "" {
		writeResponse(res, nil, result.Error)
		return
	}
	writeResponse(res, job, result.Error)
}

func HandleCreateJob(res http.ResponseWriter, req *http.Request) {
	var job model.Job
	err := json.NewDecoder(req.Body).Decode(&job)
	if ok := handlePayloadError(res, err); !ok {
		return
	}
	logger.Info.Printf("Request - create job %v\n", job)

	result := model.InsertJob(&job)
	writeResponse(res, job, result.Error)
}

func HandleUpdateJob(res http.ResponseWriter, req *http.Request) {
	var targetJob, updateJob model.Job

	id, err := uuid.Parse(mux.Vars(req)["id"])
	if ok := handleParamsError(res, err); !ok {
		return
	}
	idStr := id.String()
	targetJob = model.Job{ID: idStr}
	logger.Info.Printf("Request - update job (%s)\n", idStr)

	err = json.NewDecoder(req.Body).Decode(&updateJob)
	if ok := handlePayloadError(res, err); !ok {
		return
	}

	result := model.UpdateJob(&targetJob, &updateJob)
	writeResponse(res, targetJob, result.Error)
}

func HandleDeleteJob(res http.ResponseWriter, req *http.Request) {
	id, err := uuid.Parse(mux.Vars(req)["id"])
	if ok := handleParamsError(res, err); !ok {
		return
	}
	idStr := id.String()
	logger.Info.Printf("Request - delete job (%s)\n", idStr)

	result := model.DeleteJob(idStr)
	writeResponse(res, nil, result.Error)
}

func writeResponse(res http.ResponseWriter, data interface{}, err error) error {
	errStr := ""
	if err != nil {
		errStr = err.Error()
	}
	response := ResponsePayload{
		Data:  data,
		Error: errStr,
	}
	return json.NewEncoder(res).Encode(response)
}

func handlePayloadError(res http.ResponseWriter, err error) (ok bool) {
	ok = true
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		writeResponse(res, nil, err)
		ok = false
	}
	return
}

func handleParamsError(res http.ResponseWriter, err error) (ok bool) {
	ok = true
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		writeResponse(res, nil, err)
		ok = false
	}
	return
}
