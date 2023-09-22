package model

type Booking struct {
	Id           uint   `gorm:"primaryKey; autoIncrement:true" json:"id,omitempty"`
	BookingID    uint   `json:"booking_id,omitempty"`
	RoomID       uint   `json:"room_id,omitempty"`
	CheckInTime  string `json:"check_in_time"`
	CheckOutTime string `json:"check_out_time"`
}

type HistoryBooking struct {
	Id     uint `json:"id,omitempty"`
	UserId uint `json:"user_id,omitempty"`
}
