package main

import (
	"clientAppService/internal/config"
	"clientAppService/internal/postgresql"
	"log"
	"net/http"
	"os"
)

func main() {
	postgresClient := postgresql.NewPostgresDB(&config.DBConnectionConfig{
		User:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
		Sslmode:  os.Getenv("SLLMODE"),
	})

	err := postgresClient.Connect()

	if err != nil {
		log.Fatal("Ошибка с подключением к postgres")
	}

	log.Fatal(http.ListenAndServe(":1200", nil))

}
