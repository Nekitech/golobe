package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golobe/internal/database"
	"golobe/internal/handler"
	"golobe/internal/repository"
	services2 "golobe/internal/services"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	db := database.ConnectDB()
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")

	repos := repository.InitRepositories(db)
	services := services2.InitServices(repos)
	handlers := handler.InitHandlers(services)

	router := handlers.InitRoutes()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := router.Run(host + ":" + port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server start on port::" + host + ":" + port)

	// close database
	defer db.Close()

}
