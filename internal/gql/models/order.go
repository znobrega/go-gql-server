package models

type Order struct {
	ID           int     `json:"id" gorm:"primary key"`
	CustomerName string  `json:"customerName"`
	Amount       float64 `json:"amount"`
}

type Orders struct {
	Page  *int     `json:"page"`
	Limit *int     `json:"limit"`
	Count *int     `json:"count"`
	List  []*Order `json:"list"`
}
