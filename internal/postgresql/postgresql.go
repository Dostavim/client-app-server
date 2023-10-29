package postgresql

import (
	"clientAppService/internal/config"

	"go.uber.org/zap"
)

type PostgreSQl struct {
	cfg       *config.DBConnectionConfig
	log       *zap.Logger
	connected bool
}

func NewPostgresDB(logger *zap.Logger, cfg *config.DBConnectionConfig) *PostgreSQl {
	return &PostgreSQl{
		cfg:       cfg,
		log:       logger,
		connected: false,
	}
}
