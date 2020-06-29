package models

type Order struct {
	ID           int    `json:"id" gorm:"primary key"`
	CustomerName string `json:"costumerName"`
	OrderAmount  int    `json:"orderAmount"`
}

type Orders struct {
	Page  *int     `json:"page"`
	Limit *int     `json:"limit"`
	Count *int     `json:"count"`
	List  []*Order `json:"list"`
}
