package routes

import (
	"books-api/controller"
	"books-api/service"

	"github.com/gin-gonic/gin"
)

func BookRoute(bookService service.Service, router *gin.Engine) {
	bookHandler := controller.NewBookHandler(bookService)

	v1 := router.Group("/api/v1/")
	v1.POST("/books", bookHandler.CreateBookHandler)
	v1.GET("/books", bookHandler.GetBooksHandler)
	v1.GET("/books/:id", bookHandler.GetBookHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

}
