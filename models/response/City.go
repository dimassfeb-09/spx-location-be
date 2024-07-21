package responses

import "time"

type RequestAddCity struct {
	ID        int       `json:"id"`
	CityName  string    `validate:"required" json:"city_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
