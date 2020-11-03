package models

//Books struct
type Books struct {
	ID       string  `json:"_id"`
	BookID   string  `json:"book_id"`
	Title    string  `json:"title"`
	Owner    string  `json:"owner"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
}

//Cart struct
type Cart struct {
	ID          string  `json:"_id"`
	UserID      string  `json:"user_id"`
	Books       []Books `json:"books"`
	TotalAmount float64 `json:"total_amount"`
}
