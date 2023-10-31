package models

type Order struct {
	AuthorID      int     `json:"Author ID"`      // ID Клиента-создателя
	OrderID       int     `json:"Order ID"`       // ID Заказа, случайное число от 100_000 до 999_999
	DeliveryCost  int     `json:"Delivery Cost"`  // Цена доставки
	CategoryItem  string  `json:"Category Item"`  // Категория предмета доставки
	ItemWeight    string  `json:"Item Weight"`    // Вес заказа
	FirstPhone    string  `json:"First Phone"`    // Номер телефона для первого адреса
	FirstAddress  Address `json:"First Address"`  // Первый адрес
	SecondPhone   string  `json:"Second Phone"`   // Номер телефона для первого адреса
	SecondAddress Address `json:"Second Address"` // Второй адрес
	IsTransit     bool    `json:"Is Transit"`     // Взят ли заказ курьером
}

type Address struct {
	Street       string `json:"Street"`        // Улица
	Home         string `json:"Home"`          // Дом
	Housing      string `json:"Housing"`       // Корпус
	Entrance     string `json:"Entrance"`      // Подъезд
	Floor        string `json:"Floor"`         // Этаж
	Flat         string `json:"Flat"`          // Квартира
	IntercomCode string `json:"Intercom Code"` // Код домофона
}
