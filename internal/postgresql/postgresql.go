package postgresql

import (
	"clientAppService/internal/config"
	"clientAppService/internal/postgresql/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresConnection(cfg *config.DBConnectionConfig) (*sql.DB, error) {
	var db *sql.DB

	db, err := sql.Open("postgres", "user=postgres password=p1ecful dbname=Dostavim sslmode=disable")

	if err := db.Ping(); err != nil {
		fmt.Println("error: ", err.Error())
	}

	return db, err
}

func CreateOrder(order models.Order, db *sql.DB) {
	// Создание массивов с данными адреса, для занесения в бд
	FirstAddress := [...]string{order.FirstAddress.Street, order.FirstAddress.Home,
		order.FirstAddress.Housing, order.FirstAddress.Entrance, order.FirstAddress.Floor,
		order.FirstAddress.Flat, order.FirstAddress.IntercomCode}

	SecondAddress := [...]string{order.SecondAddress.Street, order.SecondAddress.Home,
		order.SecondAddress.Housing, order.SecondAddress.Entrance, order.SecondAddress.Floor,
		order.SecondAddress.Flat, order.SecondAddress.IntercomCode}

	// запрос к БД
	sql := `insert into Order (authorid, orderid, deliverycost, categoryitem, itemweight,
		firstphone, firstaddress, secondphone, secondaddress, istransit) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := db.Exec(sql, order.AuthorID, order.OrderID, order.DeliveryCost, order.CategoryItem,
		order.ItemWeight, order.FirstPhone, FirstAddress, order.SecondPhone,
		SecondAddress, order.IsTransit)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
