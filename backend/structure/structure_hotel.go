package structure

type Hotel struct {
	Id            uint `db:"id" gorm:"primary key; autoincrement"`
	Name          string
	Stars         uint8
	Rating        float64
	CountReviews  uint64
	Amenities     string
	Address       string
	PricePerNight uint8
	UrlImage      string
	Freebies      string
	Rooms         []Room
}
