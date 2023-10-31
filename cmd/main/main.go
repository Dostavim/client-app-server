package main

import (
	"clientAppService/internal/service"
	"clientAppService/internal/transport"
	"fmt"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var cas service.ClientAppService = &service.Service{}
	// postgresClient, err := postgresql.NewPostgresConnection(&config.DBConnectionConfig{
	//	Host: 	  os.Getenv("DB_HOST"),
	//  Port:	  os.Getenv("DB_PORT"),
	// 	User:     os.Getenv("DB_NAME"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	Dbname:   os.Getenv("DB_NAME"),
	// 	Sslmode:  os.Getenv("SLLMODE"),
	// })
	// if err != nil {
	// 	log.Fatal("Ошибка подлкючения к postgres")
	// }

	//cas := service.NewUserService(postgresClient)

	createHandler := httptransport.NewServer(
		transport.MakeCreateOrderEndpoint(cas),
		transport.DecodeCutRequest,
		transport.EncodeResponse,
	)

	http.Handle("/create", createHandler)
	fmt.Println("Server in listening on: localhost:1200/create")
	log.Fatal(http.ListenAndServe(":1200", nil))
}
