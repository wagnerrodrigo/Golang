package api

import (
	"encoding/json"
	"net/http"
	"uptime-monitor/internal/db"
)

func MonitorHandler(w http.ResponseWriter, r *http.Request) {
	results, err := db.GetAllResults()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Adicione um log para verificar os dados antes de enviar
	// log.Printf("Results: %+v", results)

	w.Header().Set("Content-Type", "application/json")
	// Codifica a resposta como JSON
	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
