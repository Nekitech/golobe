package repository

import (
	"database/sql"
	"fmt"
	"golobe/internal/database/model"
	"log"
)

type HotelDB struct {
	db *sql.DB
}

func InitHotelDB(db *sql.DB) *HotelDB {
	return &HotelDB{db: db}
}

func (repo *HotelDB) DeleteHotel(id string) error {
	query := "DELETE FROM hotels WHERE id = $1"

	_, err := repo.db.Exec(query, id)
	if err != nil {
		log.Fatalf("Unable to delete hotel: %v\n", err)
	}

	return err
}

func (repo *HotelDB) CreateHotel(hotel *model.Hotel) (*model.Hotel, error) {
	query := `
		INSERT INTO hotels (name, stars, rating, count_reviews, amenities, address, price_per_night, url_image, freebies)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`

	err := repo.db.QueryRow(query,
		&hotel.Name,
		&hotel.Stars,
		&hotel.Rating,
		&hotel.CountReviews,
		&hotel.Amenities,
		&hotel.Address,
		&hotel.PricePerNight,
		&hotel.UrlImage,
		&hotel.Freebies,
	).Scan(&hotel.Id)

	if err != nil {
		return hotel, err
	}
	return hotel, err
}

func (repo *HotelDB) GetHotels(filter *map[string]interface{}) (*[]model.Hotel, error) {

	var rows *sql.Rows
	var err error
	query := "SELECT * FROM hotels"
	if len(*filter) > 0 {

		rating, _ := (*filter)["rating"]

		fmt.Println(fmt.Sprintf("%v", rating))
		query += " WHERE rating >= " + fmt.Sprintf("%v", rating)

		rows, _ = repo.db.Query(query)

	} else {
		rows, _ = repo.db.Query(query)
	}

	var newHotels []model.Hotel

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

			return &[]model.Hotel{}, err
		}
		newHotels = append(newHotels, hotel)
	}
	return &newHotels, err
}

func (repo *HotelDB) GetHotelById(id string) (*model.Hotel, error) {
	q := `SELECT * FROM hotels WHERE id=$1;`
	row := repo.db.QueryRow(q, id)

	var hotel model.Hotel

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
		return &hotel, err
	}

	return &hotel, err
}

func (repo *HotelDB) UpdateHotel(id string, hotel *map[string]interface{}) (*model.Hotel, error) {

	query := "UPDATE hotels SET"
	params := []interface{}{}

	i := 1

	for field, value := range *hotel {
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

	result, err := repo.db.Exec(query, params...)

	// Проверка количества обновленных строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected > 0 {
		fmt.Println("Элемент успешно обновлен")

		// После обновления, можно выполнить запрос для получения обновленного элемента
		var updatedHotel model.Hotel
		err := repo.db.QueryRow("SELECT * FROM hotels WHERE id = $1", id).Scan(
			&updatedHotel.Id,
			&updatedHotel.Name,
			&updatedHotel.Stars,
			&updatedHotel.Rating,
			&updatedHotel.CountReviews,
			&updatedHotel.Amenities,
			&updatedHotel.Address,
			&updatedHotel.PricePerNight,
			&updatedHotel.UrlImage,
			&updatedHotel.Freebies,
		)
		if err != nil {
			return &model.Hotel{}, err
		}

		return &updatedHotel, err

	} else {
		fmt.Println("Элемент не найден или не обновлен")
		return &model.Hotel{}, err
	}

}
