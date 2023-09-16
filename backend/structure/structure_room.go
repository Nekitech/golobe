package structure

type Room struct {
	ID               uint `gorm:"primaryKey"`
	HotelID          uint
	Type             string
	IsViewOnCity     bool
	AmountDoubleBeds int
	AmountSingleBeds int
	CostPerNight     int
}
