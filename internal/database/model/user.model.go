package model

type User struct {
	Id          uint   `db:"id" json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Address     string `json:"address,omitempty"`
	DateOfBirth string `json:"date_of_birth,omitempty"`
}

type UserSignUp struct {
	FirstName   string `json:"first_name,omitempty" binding:"required"`
	LastName    string `json:"last_name,omitempty" binding:"required"`
	Email       string `json:"email,omitempty" binding:"required"`
	Password    string `json:"password,omitempty" binding:"required"`
	PhoneNumber string `json:"phoneNumber,omitempty" binding:"required"`
}
