package request

type EmailVerificationRequest struct {
	Mail       string `json:"email"`
	UsernamePN string `json:"username_pn"`
}
