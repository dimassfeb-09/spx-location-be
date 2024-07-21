package requests

type RequestAddSellerLocation struct {
	ID         int     `json:"id"`
	SellerName string  `json:"seller_name" binding:"required"`
	Address    string  `json:"address"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Gmaps      string  `json:"gmaps" binding:"required"`
	Status     string  `json:"status"`
	DistrictID int     `json:"district_id" binding:"required"`
	ImageURL   string  `json:"image_url" binding:"required"`
}
