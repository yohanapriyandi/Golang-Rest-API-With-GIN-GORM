package book

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Author      string      `json:"author" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required,number"`
	Discount    json.Number `json:"discount" binding:"required,number"`
}

type BookUpdate struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Price       json.Number `json:"price"`
	Author      string      `json:"author"`
	Rating      json.Number `json:"rating"`
	Discount    json.Number `json:"discount"`
}
