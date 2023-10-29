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

func NewUserService(db *postgresql.PostgreSQl) *Service {
	return &Service{
		dbClient: db,
	}
}

func (s *Service) Create() string {
	return "Create Order"
}
