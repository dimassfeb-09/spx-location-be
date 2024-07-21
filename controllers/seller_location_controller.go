package controllers

import (
	"errors"
	"fmt"
	"github.com/dimassfeb-09/spx-location-be/models"
	requests "github.com/dimassfeb-09/spx-location-be/models/request"
	responses "github.com/dimassfeb-09/spx-location-be/models/response"
	"github.com/dimassfeb-09/spx-location-be/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type SellerLocationController struct {
	*services.SellerLocationService
}

type SellerLocationControllerImpl interface {
	InsertSellerLocation(c *gin.Context)
	GetAllSellerLocation(c *gin.Context)
	GetSellerLocationBySearch(c *gin.Context)
}

func NewSellerLocationController(sellerLocation *services.SellerLocationService) SellerLocationControllerImpl {
	return &SellerLocationController{
		SellerLocationService: sellerLocation,
	}
}

func (s *SellerLocationController) InsertSellerLocation(c *gin.Context) {
	var request requests.RequestAddSellerLocation
	if err := c.ShouldBindJSON(&request); err != nil {
		var verrs gin.H
		var validationErrors validator.ValidationErrors
		fmt.Println(err)
		if errors.As(err, &validationErrors) {
			verrs = gin.H{}
			for _, v := range validationErrors {
				verrs[v.Field()] = v.ActualTag()
			}
			fmt.Println(validationErrors)
		}

		c.JSON(http.StatusBadRequest, responses.ResponseErrorWithErrors{
			Status:  400,
			Message: "Invalid request parameters",
			Errors:  verrs,
		})
		return
	}

	var sellerLocation = &models.SellerLocation{
		SellerName: request.SellerName,
		Address:    request.Address,
		Latitude:   request.Latitude,
		Longitude:  request.Longitude,
		Gmaps:      request.Gmaps,
		DistrictID: request.DistrictID,
		ImageURL:   request.ImageURL,
	}

	fmt.Println(sellerLocation)
	sellerID, err := s.SellerLocationService.InsertSellerLocation(c.Request.Context(), sellerLocation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ResponseError{
			Status:  500,
			Message: "Failed to insert seller location",
		})
		return
	}

	c.JSON(http.StatusOK, responses.ResponseSuccessWithData{
		Status:  200,
		Message: "Seller location inserted successfully",
		Data: gin.H{
			"seller_id": sellerID,
		},
	})
}

func (s *SellerLocationController) GetAllSellerLocation(c *gin.Context) {
	sellers, err := s.SellerLocationService.GetAllSellerLocation(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ResponseError{
			Status:  500,
			Message: "Failed to fetch seller locations",
		})
		return
	}

	// Return a successful response with the list of seller locations
	c.JSON(http.StatusOK, responses.ResponseSuccessWithData{
		Status:  200,
		Message: "Seller locations fetched successfully",
		Data:    sellers,
	})
}

func (s *SellerLocationController) GetSellerLocationBySearch(c *gin.Context) {

	query := c.Query("q")
	sellers, err := s.SellerLocationService.GetSellerLocationBySearch(c.Request.Context(), query)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, responses.ResponseError{
			Status:  500,
			Message: "Failed to fetch seller locations",
		})
		return
	}

	fmt.Println(query)

	c.JSON(http.StatusOK, responses.ResponseSuccessWithData{
		Status:  200,
		Message: "Successfully to retrive data seller",
		Data:    sellers,
	})
}
