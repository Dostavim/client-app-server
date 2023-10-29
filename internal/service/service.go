package service

import (
	"clientAppService/internal/postgresql"
	"context"
)

type ClientAppService interface {
	Create(ctx context.Context, order Order) string
}

type Service struct {
	dbClient *postgresql.PostgreSQl
}

type Order struct {
	FirstAddress  string `json:"FirstAddress"`
	SecondAddress string `json:"SecondAddress"`
	DeliveryItem  string `json:"Item"`
}

func NewUserService(db *postgresql.PostgreSQl) *Service {
	return &Service{
		dbClient: db,
	}
}

func (s *Service) Create(ctx context.Context, order Order) string {
	return "Create Order"
}
