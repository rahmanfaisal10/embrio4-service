package request

import "github.com/dgrijalva/jwt-go"

type LoginRequest struct {
	UsernamePN string `json:"username_pn"`
	Password   string `json:"password"`
}

type JwtCustomClaims struct {
	UserID     string `json:"user_id"`
	UsernamePN string `json:"username_pn"`
	Nama       string `json:"nama"`
	UnitKerja  string `json:"unit_kerja"`
	KodeBranch string `json:"kode_branch"`
	Jabatan    string `json:"jabatan"`
	jwt.StandardClaims
}
