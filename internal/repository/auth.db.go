package repository

import (
	"database/sql"
	"golobe/internal/database/model"
)

type AuthDB struct {
	db *sql.DB
}

func InitAuthDB(db *sql.DB) *AuthDB {
	return &AuthDB{db: db}
}

func (repo *AuthDB) CreateUser(user *model.UserSignUp) (uint, error) {
	var idUser uint

	query := `
		INSERT INTO users (first_name, last_name, email, password, phone_number)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	err := repo.db.QueryRow(query,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
	).Scan(&idUser)

	if err != nil {
		return 0, err
	}

	return idUser, err
}
