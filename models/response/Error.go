package responses

import "github.com/gin-gonic/gin"

type ResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseErrorWithErrors struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Errors  gin.H  `json:"errors"`
}
