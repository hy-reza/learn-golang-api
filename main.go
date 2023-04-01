package main

import (
	"books-api/config"
	"books-api/controller"
	"books-api/model"
	"books-api/repository"
	"books-api/service"
	"books-api/types"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getEnv(envName, fallback string) string {
	if os.Getenv(envName) != "" {
		return os.Getenv(envName)
	}
	return fallback
}

func main() {

	myEnv := types.MyEnv{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	myEnv.DB_Name = getEnv("DB_NAME", "books-api")
	myEnv.DB_User = getEnv("DB_USER", "postgres")
	myEnv.DB_Password = getEnv("DB_PASSWORD", "postgres")
	myEnv.APP_Port = getEnv("APP_PORT", ":4040")

	db := config.ConnectDB(myEnv.DB_Name, myEnv.DB_User, myEnv.DB_Password)
	db.AutoMigrate(&model.Book{})

	bookRepository := repository.NewRepository(db)
	bookService := service.NewService(bookRepository)
	bookHandler := controller.NewBookHandler(bookService)

	router := gin.Default()

	router.GET("/", controller.RootHandler)

	v1 := router.Group("/api/v1/")
	v1.POST("/books", bookHandler.CreateBookHandler)
	v1.GET("/books", bookHandler.GetBooksHandler)
	v1.GET("/books/:id", bookHandler.GetBookHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	router.Run(myEnv.APP_Port)
}
