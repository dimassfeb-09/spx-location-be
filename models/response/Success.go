package responses

type ResponseSuccess struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseSuccessWithData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
