package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/tcerqueira/tiktak/cron-backend/Nodes/Database"
)

type ResponsePayload struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func HandleGetJobsList(res http.ResponseWriter, req *http.Request) {
	jobs := []database.Job{}
	result := database.FetchAllJobs(&jobs)

	writeResponse(res, jobs, result.Error)
}

func HandleGetJob(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if ok := handleParamsError(res, err, id); !ok {
		return
	}

	job := database.Job{ID: database.JobID(id)}
	result := database.FetchJob(&job)
	if job.ID == 0 {
		writeResponse(res, nil, result.Error)
		return
	}
	writeResponse(res, job, result.Error)
}

func HandleCreateJob(res http.ResponseWriter, req *http.Request) {
	var job database.Job
	err := json.NewDecoder(req.Body).Decode(&job)
	if ok := handlePayloadError(res, err); !ok {
		return
	}

	// Do scheduled task

	result := database.InsertJob(&job)
	writeResponse(res, job, result.Error)
}

func HandleUpdateJob(res http.ResponseWriter, req *http.Request) {
	var targetJob, updateJob database.Job

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if ok := handleParamsError(res, err, id); !ok {
		return
	}
	targetJob = database.Job{ID: database.JobID(id)}

	err = json.NewDecoder(req.Body).Decode(&updateJob)
	if ok := handlePayloadError(res, err); !ok {
		return
	}

	// Update scheduled task

	result := database.UpdateJob(&targetJob, &updateJob)
	if targetJob.ID == 0 {
		writeResponse(res, nil, result.Error)
		return
	}
	writeResponse(res, targetJob, result.Error)
}

func HandleDeleteJob(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if ok := handleParamsError(res, err, id); !ok {
		return
	}

	// Delete scheduled task

	result := database.DeleteJob(database.JobID(id))
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

func handleParamsError(res http.ResponseWriter, err error, id int) (ok bool) {
	ok = true
	if err != nil || id == 0 {
		res.WriteHeader(http.StatusBadRequest)
		writeResponse(res, nil, err)
		ok = false
	}
	return
}
