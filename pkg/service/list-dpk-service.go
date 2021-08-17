package service

import (
	"rahmanfaisal10/embrio4-service/pkg/response"

	"github.com/labstack/gommon/log"
)

func (s *service) ListDpkService(mantri string) *response.BaseResponse {
	listdpk, err := s.r.ListDpkRepository(mantri)
	if err != nil {
		log.Error(err)
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	return &response.BaseResponse{
		Success: true,
		Message: "successfully to list dpk",
		Data:    listdpk,
	}
}
