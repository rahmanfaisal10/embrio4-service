package service

import (
	"rahmanfaisal10/embrio4-service/config"
	"rahmanfaisal10/embrio4-service/pkg/response"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
)

func (s *service) InsertAutoTarget(token string) *response.BaseResponse {
	//authorization
	cfg := config.Get()
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(strings.TrimPrefix(token, BEARERREMOVE), claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.SecretKey), nil
	})

	if err != nil {
		log.Error(err)
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	if claims["groups"] != "admin" {
		return &response.BaseResponse{
			Success: false,
			Message: "you do not have access to create users",
		}
	}

	err = s.r.InsertTarget()
	if err != nil {
		log.Error(err)
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	err = s.RegisterService(token)
	if err != nil {
		log.Error(err)
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	return &response.BaseResponse{
		Success: true,
		Message: "success to improve data target",
	}
}
