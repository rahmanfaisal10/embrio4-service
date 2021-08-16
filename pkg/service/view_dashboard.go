package service

import "rahmanfaisal10/embrio4-service/pkg/response"

func (s *service) ViewDashboardService(mantri string) *response.BaseResponse {
	data, err := s.r.ViewDashboard(mantri)
	if err != nil {
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	return &response.BaseResponse{
		Success: true,
		Message: "successfully view dashboard",
		Data:    data,
	}
}
