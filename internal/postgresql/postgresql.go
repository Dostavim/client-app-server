package postgresql

import (
	"clientAppService/internal/config"
	"clientAppService/internal/postgresql/models"
	"database/sql"
	"fmt"
)

func NewPostgresConnection(cfg *config.DBConnectionConfig) (*sql.DB, error) {
	var db *sql.DB

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.Dbname, cfg.Sslmode))

	return db, err
}

func Create(order models.Order, db *sql.DB) error {
	sql := `INSERT INTO Order (AuthorID, OrderID, DeliveryCost, FisrtAddress, SecondAddress,
		DeliveryItem, IsTransit) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.Exec(sql, order.AuthorID, order.OrderID, order.DeliveryCost,
		order.FirstAddress, order.SecondAddress, order.DeliveryItem, order.IsTransit)

	return err
}
