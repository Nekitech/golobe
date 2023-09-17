package model

type Room struct {
	Id               uint   `gorm:"primaryKey" json:"id,omitempty"`
	HotelID          uint   `json:"hotel_id,omitempty"`
	Type             string `json:"type,omitempty"`
	IsViewOnCity     bool   `json:"is_view_on_city,omitempty"`
	AmountDoubleBeds int    `json:"amount_double_beds,omitempty"`
	AmountSingleBeds int    `json:"amount_single_beds,omitempty"`
	CostPerNight     int    `json:"cost_per_night,omitempty"`
}
