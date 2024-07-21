package requests

import "time"

type SellerLocation struct {
	ID         int       `json:"id"`
	SellerName string    `json:"seller_name"`
	Address    string    `json:"address"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Gmaps      string    `json:"gmaps"`
	DistrictID int       `json:"district_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
