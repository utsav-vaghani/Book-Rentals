package models

//Book struct
type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	NoOfPages   int64  `json:"no_of_pages"`
	OwnerID     string `json:"owner_id"`
}
