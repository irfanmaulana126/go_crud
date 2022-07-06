package main

import (
	"belajar/book"
	"belajar/handler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")

	dsn := "root:@tcp(127.0.0.1:3306)/belajar_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed: ", err)
	}

	fmt.Println("DB connection succeeded")

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandle := handler.NewBookHandler(bookService)

	v1.POST("/book", bookHandle.CreateBooks)
	v1.GET("/books", bookHandle.GetBooks)
	v1.GET("/book/:id", bookHandle.GetByIdBook)
	v1.PUT("/book/:id", bookHandle.UpdateBook)
	v1.DELETE("/book/:id", bookHandle.DeleteBook)

	router.Run(":8888")
}

// main
// handler
// service
// repository
// db
// mysql
