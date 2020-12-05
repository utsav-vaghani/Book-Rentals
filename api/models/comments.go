package models

//Comment struct
type Comment struct {
	ID      string `json:"id"`
	BookID  string `json:"book_id" bson:"book_id"`
	UserID  string `json:"user_id" bson:"user_id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

//Comments struct
type Comments struct {
	BookID   string    `json:"book_id" bson:"book_id"`
	Comments []Comment `json:"comments"`
}
