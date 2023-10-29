package service

import (
	"clientAppService/internal/postgresql"
	"context"

	"go.uber.org/zap"
)

type ClientAppService interface {
	Create(ctx context.Context, order Order) string
}

type Service struct {
	log      *zap.Logger
	dbClient *postgresql.PostgreSQl
}

func NewUserService(logger *zap.Logger, db *postgresql.PostgreSQl) *Service {
	return &Service{
		log:      logger,
		dbClient: db,
	}
}

func (s *Service) Create() string {
	return "Create Order"
}
