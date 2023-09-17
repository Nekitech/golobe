package model

type Hotel struct {
	Id            uint    `db:"id" gorm:"primary key; autoincrement" json:"id,omitempty"`
	Name          string  `json:"name,omitempty"`
	Stars         uint8   `json:"stars,omitempty"`
	Rating        float64 `json:"rating,omitempty"`
	CountReviews  uint64  `json:"count_reviews,omitempty"`
	Amenities     string  `json:"amenities,omitempty"`
	Address       string  `json:"address,omitempty"`
	PricePerNight uint8   `json:"price_per_night,omitempty"`
	UrlImage      string  `json:"url_image,omitempty"`
	Freebies      string  `json:"freebies,omitempty"`
	Rooms         []Room  `json:"rooms,omitempty"`
}
