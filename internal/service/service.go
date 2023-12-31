package service

import (
	"clientAppService/internal/postgresql"
	"clientAppService/internal/postgresql/models"
	"context"
	"database/sql"
	"math/rand"
)

type ClientAppService interface {
	CreateOrder(ctx context.Context, AuthorID int, FirstAddress models.Address,
		SecondAddress models.Address, DeliveryItem string, ItemWeight string,
		FirstPhone string, SecondPhone string) *models.Order
}

type Service struct {
	DBClient *sql.DB
}

func NewUserService(db *sql.DB) *Service {
	return &Service{
		DBClient: db,
	}
}

func (s *Service) CreateOrder(ctx context.Context, AuthorID int, FirstAddress models.Address,
	SecondAddress models.Address, CategoryItem string, ItemWeight string, FirstPhone string,
	SecondPhone string) *models.Order {

	newOrder := &models.Order{
		AuthorID:     AuthorID,
		OrderID:      rand.Intn(100),
		DeliveryCost: 300,
		CategoryItem: CategoryItem,
		ItemWeight:   ItemWeight,
		IsTransit:    false,

		FirstPhone:  FirstPhone,
		SecondPhone: SecondPhone,

		FirstAddress:  FirstAddress,
		SecondAddress: SecondAddress,
	}

	postgresql.CreateOrder(*newOrder, s.DBClient)
	return newOrder

}
