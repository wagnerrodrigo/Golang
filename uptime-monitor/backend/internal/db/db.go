package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

// ConnectDB inicializa a conexão com o banco de dados.
func ConnectDB(connStr string) {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Verifica se a conexão está funcionando
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
}

// SaveResult insere um registro de resultado no banco de dados.
func SaveResult(url string, status bool) error {
	if db == nil {
		return fmt.Errorf("db connection is nil")
	}
	_, err := db.Exec("INSERT INTO uptime_results (url, status, checked_at) VALUES ($1, $2, NOW())", url, status)
	return err
}
