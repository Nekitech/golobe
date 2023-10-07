package model

type Hotel struct {
	Id            uint    `json:"id,omitempty"`
	Name          string  `json:"name,omitempty"`
	Stars         int     `json:"stars,omitempty"`
	Rating        float64 `json:"rating,omitempty"`
	CountReviews  int     `json:"count_reviews,omitempty"`
	Amenities     string  `json:"amenities,omitempty"`
	Address       string  `json:"address,omitempty"`
	PricePerNight int     `json:"price_per_night,omitempty"`
	UrlImage      string  `json:"url_image"`
	Freebies      string  `json:"freebies"`
}
