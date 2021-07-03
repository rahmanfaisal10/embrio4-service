package response

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Count   int64       `json:"count,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
