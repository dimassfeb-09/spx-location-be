package models

type Location struct {
	City     City     `json:"city"`
	District District `json:"district"`
}
