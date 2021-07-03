package service

import (
	"errors"
	"rahmanfaisal10/embrio4-service/config"
	"rahmanfaisal10/embrio4-service/pkg/request"
	"rahmanfaisal10/embrio4-service/pkg/response"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (s service) LoginService(req request.LoginRequest) (*response.LoginResponse, error) {
	//get user profile by username Pn
	user, err := s.r.GetUserByUsernamePN(req.UsernamePN)
	if err != nil || user == nil {
		err = errors.New("wrong username pengawai number")
		return nil, err
	}

	//checking and matching password request with password on DB
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		err = errors.New("wrong password")
		return nil, err
	}

	//set token in claims
	claims := &request.JwtCustomClaims{
		UserID:     user.UserID,
		UsernamePN: user.UsernamePN,
		Nama:       user.Nama,
		UnitKerja:  user.UnitKerja,
		KodeBranch: user.KodeBranch,
		Jabatan:    user.Jabatan,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
		},
	}

	//create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//generate encoded token and send it as response
	cfg := config.Get()
	t, err := token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return nil, err
	}

	//update last Login in table users
	err = s.r.UpdateLastLogin(*user)
	if err != nil {
		return nil, err
	}

	return &response.LoginResponse{
		ID:    int64(user.ID),
		Token: t,
	}, nil

}
