package responder

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   *string     `json:"error"`
	Data    interface{} `json:"data"`
}
