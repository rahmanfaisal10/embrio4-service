package service

import (
	"errors"
	"rahmanfaisal10/embrio4-service/config"
	"rahmanfaisal10/embrio4-service/pkg/model"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

const (
	BEARERREMOVE = "Bearer "
)

func (s *service) RegisterService(tokenReq string) error {
	//authorization
	cfg := config.Get()
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(strings.TrimPrefix(tokenReq, BEARERREMOVE), claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.SecretKey), nil
	})

	if err != nil {
		log.Error(err)
		return err
	}

	if claims["groups"] != "admin" {
		return errors.New("you do not have access to create users")
	}

	//get all mantri from table upload
	PNPengelola, err := s.r.GetAllMantriFromUpload()
	if err != nil {
		return err
	}

	for _, v := range PNPengelola {
		//checking data is exist on database
		user, _ := s.r.GetUserByUsernamePN(v.PnPengelola)
		if user != nil {
			err := errors.New("username_pn does exists, please using another username_pn")
			return err
		}

		//generate from request password
		password, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.MinCost)
		if err != nil {
			return err
		}

		newUser := &model.Users{
			ID:         0,
			UsernamePN: v.PnPengelola,
			Password:   string(password),
			Nama:       v.NamaPengelola,
			KodeBranch: v.Branch,
			LastLogin:  time.Now(),
		}

		err = s.r.CreateUser(newUser)
		if err != nil {
			return err
		}
	}

	return nil
}
