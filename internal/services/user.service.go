package services

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"golobe/internal/database/model"
	"net/http"
)

type User interface {
	AuthUser(*gin.Context)
	CreateUser(newUser *model.UserSignUp) (*uint, error)
	UpdateUser(*gin.Context)
}

type UserScheme struct {
	DB   *sql.DB
	User model.User
}

func (user *UserScheme) AuthUser(ctx *gin.Context) {

}

func (user *UserScheme) CreateUser(newUser *model.UserSignUp) (*uint, error) {

	//if err := ctx.ShouldBindJSON(&user.User); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	var idUser *uint

	query := `
		INSERT INTO users (first_name, last_name, email, password, phone_number)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	fmt.Println(newUser.FirstName)

	err := user.DB.QueryRow(query,
		newUser.FirstName,
		newUser.LastName,
		newUser.Email,
		newUser.Password,
		newUser.PhoneNumber,
	).Scan(&idUser)

	if err != nil {
		//log.Fatal().JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return nil, err
	}

	//ctx.IndentedJSON(201, user.User)
	//
	//ctx.Set("user_id", &user.User.Id)
	//ctx.Next()

	return idUser, err
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
