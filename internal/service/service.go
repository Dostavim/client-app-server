package service

import (
	"clientAppService/internal/postgresql/models"
	"context"
	"database/sql"
	"math/rand"
)

type ClientAppService interface {
	Create(ctx context.Context, AuthorID int, FirstAddress string,
		SecondAddress string, DeliveryItem string) *models.Order
}

type Service struct {
	dbClient *sql.DB
}

func NewUserService(db *sql.DB) *Service {
	return &Service{
		dbClient: db,
	}
}

func (s *Service) Create(ctx context.Context, AuthorID int, FirstAddress string,
	SecondAddress string, DeliveryItem string) *models.Order {

	newOrder := &models.Order{
		AuthorID:      AuthorID,
		OrderID:       rand.Intn(100),
		DeliveryCost:  300,
		FirstAddress:  FirstAddress,
		SecondAddress: SecondAddress,
		DeliveryItem:  DeliveryItem,
		IsTransit:     false,
	}

	//postgresql.Create(*newOrder, s.dbClient)
	return newOrder

}
