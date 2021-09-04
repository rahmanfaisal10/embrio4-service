package request

import "github.com/dgrijalva/jwt-go"

type LoginRequest struct {
	UsernamePN string `json:"username_pn"`
	Password   string `json:"password"`
}

type JwtCustomClaims struct {
	ID         string `json:"id"`
	UsernamePN string `json:"username_pn"`
	Nama       string `json:"nama"`
	Groups     string `json:"groups"`
	Branch     string `json:"branch"`
	jwt.StandardClaims
}
