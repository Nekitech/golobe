package model

type Booking struct {
	Id     uint `gorm:"primaryKey; autoIncrement:true" json:"id,omitempty"`
	UserID uint `json:"user_id,omitempty"`
	RoomID uint `json:"room_id,omitempty"`
}
