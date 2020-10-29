package dtos

//Books struct
type Books struct {
	BookID   string  `json:"book_id"`
	Title    string  `json:"title"`
	Owner    string  `json:"owner"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
}
