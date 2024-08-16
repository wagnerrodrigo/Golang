package api

import (
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/monitor", MonitorHandler).Methods("GET")
	return r
}
