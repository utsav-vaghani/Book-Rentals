package models

//Books struct
type Books struct {
	BookID   string  `json:"book_id"`
	Title    string  `json:"title"`
	Owner    string  `json:"owner"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
}

//Cart struct
type Cart struct {
	ID          string  `json:"_id" bson:"_id"`
	UserID      string  `json:"user_id" bson:"user_id"`
	Books       []Books `json:"books"`
	TotalAmount float64 `json:"total_amount" bson:"total_amount"`
}
