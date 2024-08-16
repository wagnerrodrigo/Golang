package main

import (
	"log"
	"net/http"
	"uptime-monitor/internal/api"
	"uptime-monitor/internal/config"
	"uptime-monitor/internal/db"
	"uptime-monitor/internal/monitor"

	"github.com/joho/godotenv"
)

func main() {
	// Carrega variáveis do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Configurações
	cfg := config.LoadConfig()

	// Verifique a string de conexão
	// fmt.Println("Database Connection String:", cfg.DBConnectionString)

	// Inicializa a conexão com o banco de dados
	db.ConnectDB(cfg.DBConnectionString)

	// Inicia o monitoramento
	go monitor.StartMonitoring(cfg.Urls, cfg.Interval)

	// Configura as rotas da API
	router := api.SetupRouter()

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
