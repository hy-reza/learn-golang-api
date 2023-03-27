package main

import (
	"books-api/config"
	"books-api/controller"
	"books-api/model"
	"books-api/repository"
	"books-api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	db.AutoMigrate(&model.Book{})

	bookRepository := repository.NewRepository(db)
	bookService := service.NewService(bookRepository)
	bookHandler := controller.NewBookHandler(bookService)

	router := gin.Default()
	v1 := router.Group("/api/v1/")
	v1.POST("/books", bookHandler.CreateBookHandler)
	v1.GET("/books", bookHandler.GetBooksHandler)
	v1.GET("/books/:id", bookHandler.GetBookHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	router.GET("/", controller.RootHandler)

	router.Run(":4040")
}
