package models

//Books struct
type Books struct {
	UserID   string  `json:"user_id" bson:"user_id"`
	BookID   string  `json:"book_id" bson:"book_id"`
	Title    string  `json:"title"`
	author   string  `json:"author"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
}

//Cart structs
type Cart struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id" bson:"user_id"`
	Books       []Books `json:"books"`
	TotalAmount float64 `json:"total_amount" bson:"total_amount"`
}
