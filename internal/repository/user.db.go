package repository

import (
	"database/sql"
	"fmt"
	"golobe/internal/database/model"
	"log"
)

type UserDB struct {
	db *sql.DB
}

func InitUserDB(db *sql.DB) *UserDB {
	return &UserDB{db: db}
}

func (repo *UserDB) UpdateUser(id string, updateUser *map[string]interface{}) (*model.User, error) {

	query := "UPDATE users SET"
	params := []interface{}{}

	i := 1

	for field, value := range *updateUser {
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
		var updatedUser model.User
		err := repo.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(
			&updatedUser.Id,
			&updatedUser.FirstName,
			&updatedUser.LastName,
			&updatedUser.Email,
			&updatedUser.Password,
			&updatedUser.PhoneNumber,
			&updatedUser.Address,
			&updatedUser.DateOfBirth,
		)
		if err != nil {
			return &model.User{}, err
		}

		return &updatedUser, err

	} else {
		fmt.Println("Элемент не найден или не обновлен")
		return &model.User{}, err
	}
}
