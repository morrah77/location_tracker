package domain

type OrderHistory struct {
	OrderId string `json:"order_id"`
	History []*Location `json:"history"`
}
