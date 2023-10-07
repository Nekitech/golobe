package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"golobe/internal/model"
	"log"
	"net/http"
)

type HotelScheme struct {
	DB    *sql.DB
	Hotel model.Hotel
}

func (scheme *HotelScheme) GetHotelById(ctx *gin.Context) {

	var hotel model.Hotel

	q := `SELECT * FROM hotels WHERE id=$1;`
	row := scheme.DB.QueryRow(q, ctx.Param("id"))

	err := row.Scan(
		&hotel.Id,
		&hotel.Name,
		&hotel.Stars,
		&hotel.Rating,
		&hotel.CountReviews,
		&hotel.Amenities,
		&hotel.Address,
		&hotel.PricePerNight,
		&hotel.UrlImage,
		&hotel.Freebies)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error getting hotel by id": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, &hotel)
}

func (scheme *HotelScheme) GetHotels(ctx *gin.Context) {

	var hotels []model.Hotel

	rows, err := scheme.DB.Query("SELECT * FROM hotels")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var hotel model.Hotel
		if err := rows.Scan(
			&hotel.Id,
			&hotel.Name,
			&hotel.Stars,
			&hotel.Rating,
			&hotel.CountReviews,
			&hotel.Amenities,
			&hotel.Address,
			&hotel.PricePerNight,
			&hotel.UrlImage,
			&hotel.Freebies); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error getting all hotels": err.Error()})
			return
		}
		hotels = append(hotels, hotel)
	}

	ctx.IndentedJSON(http.StatusCreated, &hotels)
}

func (scheme *HotelScheme) CreateHotel(ctx *gin.Context) {

	var hotel model.Hotel

	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO hotels (name, stars, rating, count_reviews, amenities, address, price_per_night, url_image, freebies)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`

	err := scheme.DB.QueryRow(query,
		hotel.Name,
		hotel.Stars,
		hotel.Rating,
		hotel.CountReviews,
		hotel.Amenities,
		hotel.Address,
		hotel.PricePerNight,
		hotel.UrlImage,
		hotel.Freebies,
	).Scan(&hotel.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert hotel"})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, &hotel)
}

func (scheme *HotelScheme) UpdateHotel(ctx *gin.Context) {
	id := ctx.Param("id")

	var hotel map[string]interface{}
	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE hotels SET"
	params := []interface{}{}

	i := 1

	for field, value := range hotel {
		if value != nil {
			query += " " + field + "=$" + fmt.Sprint(i) + ","
			params = append(params, value)
			i++
		}
	}

	if query[len(query)-1] == ',' {
		query = query[:len(query)-1]
	}

	query += " WHERE id=$" + fmt.Sprint(i)
	params = append(params, id)

	_, err := scheme.DB.Exec(query, params...)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update hotel"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "hotel success updated!"})

}

func (scheme *HotelScheme) DeleteHotel(ctx *gin.Context) {

	id := ctx.Param("id")

	query := "DELETE FROM hotels WHERE id = $1"

	_, err := scheme.DB.Exec(query, id)
	if err != nil {
		log.Fatalf("Unable to delete hotel: %v\n", err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Hotel deleted successfully"})
}
