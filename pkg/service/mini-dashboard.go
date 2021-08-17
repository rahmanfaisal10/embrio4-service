package service

import (
	"rahmanfaisal10/embrio4-service/pkg/response"

	"github.com/labstack/gommon/log"
)

func (service *service) ListMiniDashboardService(mantri string) *response.BaseResponse {
	miniDashboard, err := service.r.MiniDashboardRepository(mantri)
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
		Data:    miniDashboard,
	}
}
