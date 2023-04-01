package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"books-api/model"
	"books-api/service"
	"books-api/types"
)

type bookHandler struct {
	bookService service.Service
}

func NewBookHandler(bookService service.Service) *bookHandler {
	return &bookHandler{bookService}
}

func ResponseConverter(b model.Book) types.BookResponse {
	return types.BookResponse{
		ID:     b.ID,
		Author: b.Author,
		Desc:   b.Desc,
		Title:  b.Title,
	}
}

func (h *bookHandler) CreateBookHandler(ctx *gin.Context) {
	var bookRequest types.BookRequest

	err := ctx.ShouldBindJSON(&bookRequest)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	newBook, err := h.bookService.Create(bookRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := ResponseConverter(newBook)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"msg":    "successfully created a book !",
		"data":   bookResponse,
	})
}

func (h *bookHandler) UpdateBookHandler(ctx *gin.Context) {
	var bookRequest types.BookRequest

	idString := ctx.Param("id")

	id, _ := strconv.Atoi(idString)
	err := ctx.ShouldBindJSON(&bookRequest)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	newBook, err := h.bookService.Update(id, bookRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := ResponseConverter(newBook)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"msg":    "successfully updated a book !",
		"data":   bookResponse,
	})
}

func (h *bookHandler) GetBooksHandler(ctx *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []types.BookResponse

	for _, b := range books {
		bookResponse := ResponseConverter(b)

		booksResponse = append(booksResponse, bookResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":   booksResponse,
		"msg":    "successfully get entire book !",
		"status": "success",
	})
}

func (h *bookHandler) GetBookHandler(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.FindByID(int(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := ResponseConverter(book)

	ctx.JSON(http.StatusOK, gin.H{
		"data":   bookResponse,
		"msg":    "successfully get a book !",
		"status": "success",
	})
}

func (h *bookHandler) DeleteBookHandler(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.Delete(int(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := ResponseConverter(book)

	ctx.JSON(http.StatusOK, gin.H{
		"data":   bookResponse,
		"msg":    "successfully delete a book !",
		"status": "success",
	})
}
