package model

type User struct {
	Id             uint           `db:"id" gorm:"primary key; autoincrement" json:"id,omitempty"`
	FirstName      string         `json:"first_name,omitempty"`
	LastName       string         `json:"last_name,omitempty"`
	Email          string         `json:"email,omitempty"`
	Password       string         `json:"password,omitempty"`
	PhoneNumber    string         `json:"phone_number,omitempty"`
	Address        string         `json:"address,omitempty"`
	DateOfBirth    string         `json:"date_of_birth,omitempty"`
	HistoryBooking HistoryBooking `json:"history_booking" gorm:"foreignKey:UserId"`
}
