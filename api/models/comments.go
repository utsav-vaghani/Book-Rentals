package models

//Comment struct
type Comment struct {
	ID      string `json:"_id" bson:"_id"`
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

//Comments struct
type Comments struct {
	ID       string    `json:"_id" bson:"_id"`
	BookID   string    `json:"book_id"`
	Comments []Comment `json:"comments"`
}
