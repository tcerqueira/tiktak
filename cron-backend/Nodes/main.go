package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	h "github.com/tcerqueira/tiktak/cron-backend/Nodes/Handlers"
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
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"*"})
	origins := handlers.AllowedOrigins([]string{"*"})
	// List all crons
	router.HandleFunc("/cron", h.HandleGetJobsList).Methods("GET")
	// Fetch cron with id
	router.HandleFunc("/cron/{id}", h.HandleGetJob).Methods("GET")
	// Create cron
	router.HandleFunc("/cron", h.HandleCreateJob).Methods("POST")
	// Update cron
	router.HandleFunc("/cron/{id}", h.HandleUpdateJob).Methods("PUT")
	// Delete cron
	router.HandleFunc("/cron/{id}", h.HandleDeleteJob).Methods("DELETE")

	// Start server
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":"+sv_port, handlers.CORS(credentials, methods, origins)(router)))
}
