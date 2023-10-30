package models

type Order struct {
	AuthorID     int `json:"Author ID"`
	OrderID      int `json:"Order ID"`
	DeliveryCost int `json:"Delivery Cost"`

	FirstAddress  string `json:"First Address"`
	SecondAddress string `json:"Second Address"`
	DeliveryItem  string `json:"Delivery Item"`

	IsTransit bool `json:"Is Transit"`
}
