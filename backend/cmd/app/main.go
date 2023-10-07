package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golobe/internal/database"
	routes2 "golobe/internal/routes"
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

	router := gin.Default()

	routes2.BookingRoute(db, router)
	routes2.HotelRoute(db, router)
	routes2.RoomRoute(db, router)
	routes2.UserRoute(db, router)

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
