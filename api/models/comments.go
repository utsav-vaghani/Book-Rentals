package models

//Comment struct
type Comment struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

//Comments struct
type Comments struct {
	BookID   string    `json:"book_id"`
	Comments []Comment `json:"comments"`
}
