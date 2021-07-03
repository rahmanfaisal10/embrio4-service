package response

type LoginResponse struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
}
