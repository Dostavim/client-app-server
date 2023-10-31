package main

import (
	"clientAppService/internal/config"
	"clientAppService/internal/postgresql"
	"clientAppService/internal/service"
	"clientAppService/internal/transport"
	"fmt"
	"log"
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	postgresClient, err := postgresql.NewPostgresConnection(&config.DBConnectionConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
	})
	if err != nil {
		log.Fatal("Ошибка подлкючения к postgres", err)
	}

	var cas service.ClientAppService = &service.Service{
		DBClient: postgresClient,
	}

	createHandler := httptransport.NewServer(
		transport.MakeCreateOrderEndpoint(cas),
		transport.DecodeCreateOrderRequest,
		transport.EncodeResponse,
	)

	http.Handle("/create", createHandler)
	fmt.Println("Server in listening on: localhost:1200/create")
	log.Fatal(http.ListenAndServe(":1200", nil))
}
