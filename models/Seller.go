package models

type SellerLocation struct {
	ID            int           `json:"id"`
	SellerName    string        `json:"seller_name"`
	Address       string        `json:"address"`
	Latitude      float64       `json:"latitude"`
	Longitude     float64       `json:"longitude"`
	Gmaps         string        `json:"gmaps"`
	DistrictID    int           `json:"district_id"`
	ImageURL      string        `json:"image_url"`
	TimeOperation TimeOperation `json:"time_operation"`
}

type DetailSeller struct {
	SellerLocation SellerLocation `json:"seller_info"`
	Location       Location       `json:"location"`
}
