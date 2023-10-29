package postgresql

import (
	"clientAppService/internal/config"
	"database/sql"
	"fmt"
)

type PostgreSQl struct {
	cfg       *config.DBConnectionConfig
	connected bool
}

func NewPostgresDB(cfg *config.DBConnectionConfig) *PostgreSQl {
	return &PostgreSQl{
		cfg:       cfg,
		connected: false,
	}
}

func (pg *PostgreSQl) Connect() error {
	_, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		pg.cfg.User, pg.cfg.Password, pg.cfg.Dbname, pg.cfg.Sslmode))

	return err
}
