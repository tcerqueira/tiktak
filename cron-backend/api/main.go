package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	h "github.com/tcerqueira/tiktak/cron-backend/api/handlers"
	logger "github.com/tcerqueira/tiktak/cron-backend/internal/logger"
)

var (
	svPort string
)

func init() {
	if len(os.Args) > 1 {
		svPort = os.Args[1]
	} else {
		svPort = "8080"
	}
}

func main() {
	router := mux.NewRouter()
	// headers := handlers.AllowedHeaders([]string{"*"})
	// credentials := handlers.AllowCredentials()
	// methods := handlers.AllowedMethods([]string{"*"})
	// origins := handlers.AllowedOrigins([]string{"*"})
	// List all crons
	// router.HandleFunc("/cron", h.HandleGetJobsList).Methods("GET", "OPTIONS")
	router.HandleFunc("/cron", WithCors(h.HandleGetJobsList)).Methods("GET", "OPTIONS")
	// Fetch cron with id
	// router.HandleFunc("/cron/{id}", h.HandleGetJob).Methods("GET", "OPTIONS")
	router.HandleFunc("/cron/{id}", WithCors(h.HandleGetJob)).Methods("GET", "OPTIONS")
	// Create cron
	// router.HandleFunc("/cron", h.HandleCreateJob).Methods("POST", "OPTIONS")
	router.HandleFunc("/cron", WithCors(h.HandleCreateJob)).Methods("POST", "OPTIONS")
	// Update cron
	// router.HandleFunc("/cron/{id}", h.HandleUpdateJob).Methods("PUT", "OPTIONS")
	router.HandleFunc("/cron/{id}", WithCors(h.HandleUpdateJob)).Methods("PUT", "OPTIONS")
	// Delete cron
	// router.HandleFunc("/cron/{id}", h.HandleDeleteJob).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/cron/{id}", WithCors(h.HandleDeleteJob)).Methods("DELETE", "OPTIONS")

	// Migrate cron jobs already in DB
	// cronjob.Migrate([]Jobs{})

	// Start server
	logger.Info.Printf("Starting server. Listening on PORT %s\n", svPort)
	log.Fatal(http.ListenAndServe(":"+svPort, router))
	// log.Fatal(http.ListenAndServe(":"+svPort, handlers.CORS(headers, credentials, methods, origins)(router)))
}

func WithCors(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Headers", "*")
		res.Header().Set("Access-Control-Allow-Credentials", "true")
		res.Header().Set("Access-Control-Allow-Methods", "*")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		handler(res, req)
	}
}
