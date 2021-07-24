package service

import (
	"errors"
	"rahmanfaisal10/embrio4-service/pkg/request"
	"rahmanfaisal10/embrio4-service/pkg/response"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) ChangedPasswordService(request request.ChangePasswordRequest) *response.BaseResponse {
	//get user profile by username Pn
	user, err := s.r.GetUserByUsernamePN(request.UsernamePN)
	if err != nil || user == nil {
		log.Error(err)
		err = errors.New("wrong username pengawai number")
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	//checking and matching password request with password on DB
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		err = errors.New("your password is wrong")
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	if request.Password == request.NewPassword {
		return &response.BaseResponse{
			Success: false,
			Message: "dont use password the same with previous password",
		}
	}

	//update to db
	password, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.MinCost)
	if err != nil {
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	user.Password = string(password)

	err = s.r.UpdatePassword(*user)
	if err != nil {
		log.Error(err)
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	return &response.BaseResponse{
		Success: true,
		Message: "successfully for change password and please login again",
	}
}
