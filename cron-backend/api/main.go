package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	h "github.com/tcerqueira/tiktak/cron-backend/api/handlers"
	logger "github.com/tcerqueira/tiktak/cron-backend/api/logger"
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

	// Migrate cron jobs already in DB
	// cronjob.Migrate([]Jobs{})
	id, err := uuid.Parse("00000000-0000-0000-0000-000000000000")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)

	// Start server
	logger.Info.Printf("Starting server. Listening on PORT %s\n", sv_port)
	log.Fatal(http.ListenAndServe(":"+sv_port, handlers.CORS(credentials, methods, origins)(router)))
}
