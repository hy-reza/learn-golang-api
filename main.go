package main

import (
	"books-api/config"
	"books-api/controller"
	"books-api/model"
	"books-api/repository"
	"books-api/routes"
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

	server := gin.Default()
	server.GET("/", controller.RootHandler)

	routes.BookRoute(bookService, server)

	server.Run(myEnv.APP_Port)
}
