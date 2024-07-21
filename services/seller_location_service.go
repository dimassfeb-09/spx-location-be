package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dimassfeb-09/spx-location-be/models"
	requests "github.com/dimassfeb-09/spx-location-be/models/request"
	responses "github.com/dimassfeb-09/spx-location-be/models/response"
)

type SellerLocationService struct {
	DB *sql.DB
}

type SellerLocationServiceImpl interface {
	InsertSellerLocation(ctx context.Context, sellerLocation *models.SellerLocation) (sellerID int, err error)
	UpdateSellerLocation(ctx context.Context, seller *requests.RequestAddSellerLocation) (*responses.ResponseSuccess, error)
	DeleteSellerLocation(ctx context.Context, seller *requests.RequestAddSellerLocation) (*responses.ResponseSuccess, error)
	GetAllSellerLocation(ctx context.Context) ([]models.SellerLocation, error)
	GetSellerLocationBySearch(ctx context.Context, query string) ([]models.DetailSeller, error)
}

func NewSellerLocationService(DB *sql.DB) *SellerLocationService {
	return &SellerLocationService{
		DB: DB,
	}
}

func (s *SellerLocationService) InsertSellerLocation(ctx context.Context, sellerLocation *models.SellerLocation) (sellerID int, err error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	query := `INSERT INTO request_add_seller_location (seller_name, address, latitude, longitude, gmaps, district_id, image_url) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	var id int
	err = tx.QueryRowContext(ctx, query,
		sellerLocation.SellerName,
		sql.NullString{String: sellerLocation.Address, Valid: sellerLocation.Address != ""},
		sellerLocation.Latitude,
		sellerLocation.Longitude,
		sql.NullString{String: sellerLocation.Gmaps, Valid: sellerLocation.Gmaps != ""},
		sellerLocation.DistrictID,
		sellerLocation.ImageURL).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SellerLocationService) UpdateSellerLocation(ctx context.Context, seller *requests.RequestAddSellerLocation) (*responses.ResponseSuccess, error) {
	// TODO: Implement the update logic
	return nil, fmt.Errorf("update seller location not implemented")
}

func (s *SellerLocationService) DeleteSellerLocation(ctx context.Context, seller *requests.RequestAddSellerLocation) (*responses.ResponseSuccess, error) {
	// TODO: Implement the delete logic
	return nil, fmt.Errorf("delete seller location not implemented")
}

func (s *SellerLocationService) GetAllSellerLocation(ctx context.Context) ([]models.SellerLocation, error) {
	query := `SELECT id, seller_name, address, latitude, longitude, gmaps, district_id, created_at, updated_at FROM seller_location`

	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sellerLocations []models.SellerLocation
	for rows.Next() {
		var loc models.SellerLocation
		var updatedAt sql.NullTime
		err := rows.Scan(
			&loc.ID,
			&loc.SellerName,
			&loc.Address,
			&loc.Latitude,
			&loc.Longitude,
			&loc.Gmaps,
			&loc.DistrictID,
			&loc.TimeOperation.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		if updatedAt.Valid {
			loc.TimeOperation.UpdatedAt = updatedAt.Time
		}
		sellerLocations = append(sellerLocations, loc)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}

	return sellerLocations, nil
}

func (s *SellerLocationService) GetSellerLocationBySearch(ctx context.Context, searchQuery string) ([]models.DetailSeller, error) {
	query := `
	SELECT 
		sl.id AS seller_id,
		sl.seller_name,
		sl.address,
		sl.gmaps,
		sl.latitude,
		sl.longitude,
		d.id AS district_id,
		d.district_name,
		c.id AS city_id,
		c.city_name,
		sl.created_at,
		sl.updated_at
	FROM 
		seller_location sl
	JOIN 
		district d ON sl.district_id = d.id
	JOIN 
		city c ON d.city_id = c.id
	WHERE 
		c.city_name ILIKE $1 OR d.district_name ILIKE $2;
	`

	searchPattern := "%" + searchQuery + "%"

	rows, err := s.DB.QueryContext(ctx, query, searchPattern, searchPattern)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var detailSellers []models.DetailSeller
	for rows.Next() {
		var detailSeller models.DetailSeller
		var sl models.SellerLocation
		var location models.Location
		var updatedAt sql.NullTime

		err := rows.Scan(
			&sl.ID,
			&sl.SellerName,
			&sl.Address,
			&sl.Gmaps,
			&sl.Latitude,
			&sl.Longitude,
			&location.District.ID,
			&location.District.DistrictName,
			&location.City.ID,
			&location.City.CityName,
			&sl.TimeOperation.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		if updatedAt.Valid {
			sl.TimeOperation.UpdatedAt = updatedAt.Time
		}

		detailSeller.SellerLocation = sl
		detailSeller.Location = location
		detailSellers = append(detailSellers, detailSeller)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}

	return detailSellers, nil
}
