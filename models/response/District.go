package responses

import "time"

type RequestDistrict struct {
	ID           int       `json:"id"`
	DistrictName string    `json:"district_name"`
	CityID       int       `json:"city_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
