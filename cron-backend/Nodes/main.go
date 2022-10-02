package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	handlers "github.com/tcerqueira/tiktak/cron-backend/Nodes/Handlers"
)

var (
	sv_port string
)

func init() {
	if len(os.Args) > 1 {
		sv_port = os.Args[1]
	} else {
		sv_port = "8080"
	}
}

func main() {
	router := mux.NewRouter()
	// List all crons
	router.HandleFunc("/cron", handlers.HandleGetJobsList).Methods("GET")
	// Fetch cron with id
	router.HandleFunc("/cron/{id}", handlers.HandleGetJob).Methods("GET")
	// Create cron
	router.HandleFunc("/cron", handlers.HandleCreateJob).Methods("POST")
	// Update cron
	router.HandleFunc("/cron/{id}", handlers.HandleUpdateJob).Methods("PUT")
	// Delete cron
	router.HandleFunc("/cron/{id}", handlers.HandleDeleteJob).Methods("DELETE")

	// Start server
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":"+sv_port, router))
}
