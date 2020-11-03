package models

import "time"

//StatusType types of order status
type StatusType string

const (
	Pending   StatusType = "Pending"
	Placed    StatusType = "Placed"
	Delivered StatusType = "Delivered"
)

//Order struct
type Order struct {
	ID          string     `json:"_id"`
	UserID      string     `json:"user_id"`
	Books       []Books    `json:"books"`
	TotalAmount float64    `json:"total_amount"`
	Time        time.Time  `json:"time"`
	Status      StatusType `json:"status"`
}
