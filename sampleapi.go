package main

import (
	"io"
	"net/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func Sample(w http.ResponseWriter, r *http.Request) {
	log.Info("API Sample is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"Running": true}`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	log.Info("Starting Sample API server")
	router := mux.NewRouter()
	router.HandleFunc("/sample", Sample).Methods("GET")
	http.ListenAndServe(":8000", router)
}