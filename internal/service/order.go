package service

type Order struct {
	AuthorID      int    `json:"AuthorId"`
	FirstAddress  string `json:"FirstAddress"`
	SecondAddress string `json:"SecondAddress"`
	DeliveryItem  string `json:"Item"`
}
