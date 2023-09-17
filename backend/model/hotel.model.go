package model

type Hotel struct {
	Id            uint    `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true"`
	Name          string  `json:"name,omitempty"`
	Stars         int     `json:"stars,omitempty"`
	Rating        float64 `json:"rating,omitempty"`
	CountReviews  int     `json:"count_reviews,omitempty"`
	Amenities     string  `json:"amenities,omitempty"`
	Address       string  `json:"address,omitempty"`
	PricePerNight int     `json:"price_per_night,omitempty"`
	UrlImage      string  `json:"url_image,omitempty"`
	Freebies      string  `json:"freebies,omitempty"`
	Rooms         []Room  `gorm:"foreignKey:HotelId;constraint:OnDelete:CASCADE" json:"rooms,omitempty"`
}
