package responses

import "github.com/dimassfeb-09/spx-location-be/models"

type SellerLocationSearch struct {
	Seller   models.SellerLocation
	District models.District
	City     models.City
}
