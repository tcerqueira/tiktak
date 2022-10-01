package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/tcerqueira/tiktak/cron-backend/Dispatcher/pkg/Handlers"
)

func main() {
	router := mux.NewRouter()
	// List all jobs
	router.HandleFunc("/job", handlers.HandleGetJobsList).Methods("GET")
	// Fetch job with id
	router.HandleFunc("/job/{id}", handlers.HandleGetJob).Methods("GET")
	// Create job
	router.HandleFunc("/job", handlers.HandleCreateJob).Methods("POST")
	// Update job
	router.HandleFunc("/job/{id}", handlers.HandleUpdateJob).Methods("PUT")
	// Delete job
	router.HandleFunc("/job/{id}", handlers.HandleDeleteJob).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
