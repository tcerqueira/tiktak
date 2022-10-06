package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	logger "github.com/tcerqueira/tiktak/cron-backend/internal/logger"
	"gorm.io/gorm/logger"
)

var (
	sv_port string
)

func init() {
	if len(os.Args) > 1 {
		sv_port = os.Args[1]
	} else {
		sv_port = "8050"
	}
}

func main() {
	router := mux.NewRouter()
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"*"})
	origins := handlers.AllowedOrigins([]string{"*"})
	// Log requests
	router.HandleFunc("/webhook", handleGet).Methods("GET")
	router.HandleFunc("/webhook", handleOthers).Methods("POST", "PUT", "PATCH", "DELETE")

	// Start server
	logger.Info.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":"+sv_port, handlers.CORS(credentials, methods, origins)(router)))
}

type RequestBody struct {
	Body string
}

func handleGet(res http.ResponseWriter, req *http.Request) {
	body := req.URL.Query().Get("body")
	handle(res, req, body)
}

func handleOthers(res http.ResponseWriter, req *http.Request) {
	var body RequestBody
	json.NewDecoder(req.Body).Decode(&body)
	handle(res, req, body.Body)
}

func handle(res http.ResponseWriter, req *http.Request, body string) {
	fmt.Printf("%s - %s\n", time.Now().Format("2006-02-01 15:04:05"), body)
	res.Write([]byte(""))
}
