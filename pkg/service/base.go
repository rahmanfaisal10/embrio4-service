package service

import (
	"rahmanfaisal10/embrio4-service/pkg/repository"
	"rahmanfaisal10/embrio4-service/pkg/request"
	"rahmanfaisal10/embrio4-service/pkg/response"
)

type Service interface {
	LoginService(req request.LoginRequest) (*response.LoginResponse, error)
	RegisterService(request request.RegisterRequest, tokenString string) error
	EmailVerificationService(request request.EmailVerificationRequest) *response.BaseResponse
	ChangedPasswordService(request request.ChangePasswordRequest) *response.BaseResponse
}

type service struct {
	r repository.Repository
}

func InitService(r repository.Repository) Service {
	return &service{r}
}
