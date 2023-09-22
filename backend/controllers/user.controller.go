package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"golobe/model"
	"net/http"
)

type UserScheme struct {
	DB   *sql.DB
	User model.User
}

func (user *UserScheme) CreateUser(ctx *gin.Context) {

	if err := ctx.ShouldBindJSON(&user.User); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO users (first_name, last_name, email, password, phone_number, address, date_of_birth)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	err := user.DB.QueryRow(query,
		&user.User.FirstName,
		&user.User.LastName,
		&user.User.Email,
		&user.User.Password,
		&user.User.PhoneNumber,
		&user.User.Address,
		&user.User.DateOfBirth,
	).Scan(&user.User.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.IndentedJSON(201, user.User)

	ctx.Set("user_id", &user.User.Id)
	ctx.Next()
}

func (user *UserScheme) UpdateUser(ctx *gin.Context) {

	id := ctx.Param("id")

	var updateUser map[string]interface{}

	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE users SET"
	params := []interface{}{}

	i := 1

	for field, value := range updateUser {
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

	_, err := user.DB.Exec(query, params...)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.IndentedJSON(201, user.User)

}
