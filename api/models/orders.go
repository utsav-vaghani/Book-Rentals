package models

import (
	"../dtos"
)

//StatusType types of order status
type StatusType string

const (
	Pending   StatusType = "Pending"
	Placed    StatusType = "Placed"
	Delivered StatusType = "Delivered"
)

//Orders struct
type Orders struct {
	UserID      string       `json:"user_id"`
	Books       []dtos.Books `json:"books"`
	TotalAmount float64      `json:"total_amount"`
	Time        int64        `json:"time"`
	Status      StatusType   `json:"status"`
}
