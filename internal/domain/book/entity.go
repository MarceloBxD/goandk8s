package book

type Book struct {
	ID uint `json: "book_id"`
	Title string `json: "title"`
	Author string `json: "author"`
	Price float64 `json: "price"`
}