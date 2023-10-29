package main

import (
	"clientAppService/internal/config"
	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {
	postgresClient := &config.DBConnectionConfig{
		User:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
		Sslmode:  os.Getenv("SLLMODE"),
	}

	_, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		postgresClient.User, postgresClient.Password, postgresClient.Dbname, postgresClient.Sslmode))

	if err != nil {
		log.Fatal(err)
	}
}
