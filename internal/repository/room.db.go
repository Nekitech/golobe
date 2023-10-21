package repository

import (
	"database/sql"
	"fmt"
	"golobe/internal/database/model"
	"log"
)

type RoomDB struct {
	db *sql.DB
}

func InitRoomDB(db *sql.DB) *RoomDB {
	return &RoomDB{db: db}
}

func (repo *RoomDB) CreateRoom(room *model.Room) (*model.Room, error) {
	query := `
		INSERT INTO rooms (hotel_id, type, is_view_on_city, amount_double_beds, amount_single_beds, cost_per_night) 
		values ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	err := repo.db.QueryRow(query,
		&room.HotelId,
		&room.Type,
		&room.IsViewOnCity,
		&room.AmountDoubleBeds,
		&room.AmountSingleBeds,
		&room.CostPerNight,
	).Scan(&room.Id)

	if err != nil {
		return room, err
	}
	return room, err

}

func (repo *RoomDB) UpdateRoom(id string, room *map[string]interface{}) (*model.Room, error) {

	query := "UPDATE rooms SET"
	params := []interface{}{}

	i := 1

	for field, value := range *room {
		if value != nil {
			query += " " + field + "=$" + fmt.Sprint(i) + ","
			params = append(params, value)
			i++
		}
	}

	fmt.Println(query)

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
		var updatedRoom model.Room
		err := repo.db.QueryRow("SELECT * FROM rooms WHERE id = $1", id).Scan(
			&updatedRoom.Id,
			&updatedRoom.HotelId,
			&updatedRoom.Type,
			&updatedRoom.IsViewOnCity,
			&updatedRoom.AmountDoubleBeds,
			&updatedRoom.AmountSingleBeds,
			&updatedRoom.CostPerNight,
		)
		if err != nil {
			return &model.Room{}, err
		}

		return &updatedRoom, err

	} else {
		fmt.Println("Элемент не найден или не обновлен")
		return &model.Room{}, err
	}
}
