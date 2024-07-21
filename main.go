package main

import (
	"context"
	"fmt"
	"github.com/dimassfeb-09/spx-location-be/config"
	"github.com/dimassfeb-09/spx-location-be/controllers"
	"github.com/dimassfeb-09/spx-location-be/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func uploadHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to retrieve file"})
		return
	}
	defer file.Close()

	url, err := services.UploadImageURLOrder(context.Background(), file, "uploaded-file-name")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully!", "url": url})
}

func main() {

	db := config.ConnectDB()

	ctx := context.Background()
	if err := services.InitFirebaseApp(ctx, "serviceAccountKey.json"); err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	r := gin.Default()

	sellerLocationService := services.NewSellerLocationService(db)
	sellerLocationController := controllers.NewSellerLocationController(sellerLocationService)

	r.POST("/seller/upload", uploadHandler)
	r.POST("/seller", sellerLocationController.InsertSellerLocation)
	r.GET("/seller", sellerLocationController.GetAllSellerLocation)
	r.GET("/seller/search", sellerLocationController.GetSellerLocationBySearch)
	err := r.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
