package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golobe/database"
	"golobe/routes"
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

	routes.BookingRoute(db, router)
	routes.HotelRoute(db, router)
	routes.RoomRoute(db, router)
	routes.UserRoute(db, router)

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
