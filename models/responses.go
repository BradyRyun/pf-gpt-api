package models

func CreateResponse(err error, message string, data interface{}) ApiResponse {
	var e string
	if err != nil {
		e = err.Error()
	}
	return ApiResponse{
		Error:   e,
		Message: message,
		Data:    data,
	}
}

type ApiResponse struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}
