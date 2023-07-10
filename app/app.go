package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/time", getCurrentTime).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", router))
}
