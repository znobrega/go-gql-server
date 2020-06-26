package models

type Order struct {
	ID           int    `json:"id" gorm:"primary key"`
	CustomerName string `json:"costumerName"`
	OrderAmount  int    `json:"orderAmount"`
}
