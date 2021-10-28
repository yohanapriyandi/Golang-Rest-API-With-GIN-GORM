package book

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	Discount    int    `json:"discount"`
}
