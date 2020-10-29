package models

import "../dtos"

//Cart struct
type Cart struct {
	UserID      string       `json:"user_id"`
	Books       []dtos.Books `json:"books"`
	TotalAmount float64      `json:"total_amount"`
}
