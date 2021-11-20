package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Koneksi ke database mysql
	dsn := "root:1tc5_45uNn41-1@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// bookRequest := book.BookRequest{
	// 	Title: "Fundamental Goroutines",
	// 	Price: "250000",
	// }

	// bookService.Create(bookRequest)

	/*=========================================================================================*/

	// books, err := bookRepository.FindAll()
	// for _, book := range books {
	// 	fmt.Println("Title:", book.Title)
	// }

	/*=========================================================================================*/

	// book, err := bookRepository.FindByID(3)
	// fmt.Println("Title :", book.Title)

	/*=========================================================================================*/
	// book := book.Book{
	// 	Title:       "Active Storage Ruby  on Rails",
	// 	Description: "How to use active storage for save file",
	// 	Author:      "Yohan Apriyandi",
	// 	Price:       100000,
	// 	Discount:    0,
	// 	Rating:      3,
	// }

	// bookRepository.Create(book)
	/*=========================================================================================*/

	/*Create Data*/
	// CRUD
	// book := book.Book{}
	// book.Title = "Membangun Aplikasi Android Untuk Pemula"
	// book.Description = "Kupas tuntas bagaimana membuat aplikasi android bagi pemula dengan studi kasu real project"
	// book.Price = 150000
	// book.Discount = 10
	// book.Author = "Yohan Apriyandi"
	// book.Rating = 5

	// err = db.Debug().Create(&book).Error
	// if err != nil {
	// 	fmt.Println("=======================================")
	// 	fmt.Print("Error creating book record")
	// 	fmt.Println("=======================================")
	// }

	/*GET Specific Data*/
	// var book book.Book
	// err = db.Debug().First(&book, 3).Error

	// if err != nil {
	// 	fmt.Println("=======================================")
	// 	fmt.Print("Error finding book record")
	// 	fmt.Println("=======================================")
	// }
	// fmt.Println("Title :", book.Title)
	// fmt.Println("book object %v", book)

	/*GET All Data*/

	// var books []book.Book
	// err = db.Debug().Find(&books).Error

	// if err != nil {
	// 	fmt.Println("=======================================")
	// 	fmt.Print("Error finding book record")
	// 	fmt.Println("=======================================")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title :", b.Title)
	// 	fmt.Println("Description :", b.Description)
	// 	fmt.Println("Author :", b.Author)
	// 	// fmt.Println("book object %v", b)
	// }

	/*Get with filter*/
	// var books []book.Book
	// err = db.Debug().Where("rating = ?", 5).Find(&books).Error

	// if err != nil {
	// 	fmt.Println("=======================================")
	// 	fmt.Print("Error finding book record")
	// 	fmt.Println("=======================================")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title :", b.Title)
	// 	fmt.Println("Description :", b.Description)
	// 	fmt.Println("Author :", b.Author)
	// 	// fmt.Println("book object %v", b)
	// }

	// /*Update data*/
	// var book book.Book
	// err = db.Debug().Where("id = ?", 2).First(&book).Error

	// if err != nil {
	// 	fmt.Println("=======================================")
	// 	fmt.Print("Error finding book record")
	// 	fmt.Println("=======================================")
	// }

	// book.Title = "Membangun web aplikasi dengan framework Ruby on Rails 6 & Postgresql"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("=======================================")
	// 	fmt.Print("Error updating book record")
	// 	fmt.Println("=======================================")
	// }

	/*Delete data*/
	// var book book.Book
	// err = db.Debug().Where("id = ?", 4).First(&book).Error

	// if err != nil {
	// 	fmt.Println("=======================================")
	// 	fmt.Print("Error finding book record")
	// 	fmt.Println("=======================================")
	// }

	// err = db.Debug().Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("=======================================")
	// 	fmt.Print("Error deleting book record")
	// 	fmt.Println("=======================================")
	// }

	// initialisasi var router
	router := gin.Default()

	/*
		this for versioning API
		you can add your versioning API below
	*/

	v1 := router.Group("/api/v1")

	// GET router
	v1.GET("/books", bookHandler.GetBooksHandler)
	v1.GET("/books/:id", bookHandler.GetBookHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBookHandler)

	// POST router
	v1.POST("/books", bookHandler.CreateBookHandler)

	// DELETE router
	v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	router.Run()

}
