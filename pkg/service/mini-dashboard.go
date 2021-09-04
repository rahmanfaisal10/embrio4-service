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

	getAllJatuhTempo, err := service.r.GetBelumJatuhTempo(mantri)
	if err != nil {
		log.Error(err)
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	minidashboardResponse := new(response.MiniDashboardsResponse)

	for _, v := range miniDashboard {
		switch v.Status {
		case "janji setor":
			minidashboardResponse.DpkJanjiStor = v.DpkTotal
			minidashboardResponse.CountDpkJanjiStor = v.Count
		case "tidak bayar":
			minidashboardResponse.DpkTidakBayar = v.DpkTotal
			minidashboardResponse.CountDpkTidakBayar = v.Count
		case "proses ob/agf":
			minidashboardResponse.DpkObAgf = v.DpkTotal
			minidashboardResponse.CountDpkObAgf = v.Count
		case "":
			minidashboardResponse.PosisiDPK = v.DpkTotal
			minidashboardResponse.CountPosisiDPK = v.Count
		}
	}

	minidashboardResponse.DpkBelumJatuhTempo = getAllJatuhTempo.DpkTotal
	minidashboardResponse.CountDpkBelumJatuhTempo = getAllJatuhTempo.Count
	minidashboardResponse.PrognosaDpk = minidashboardResponse.PosisiDPK - minidashboardResponse.DpkJanjiStor - minidashboardResponse.DpkObAgf + minidashboardResponse.DpkTidakBayar + minidashboardResponse.DpkBelumJatuhTempo
	minidashboardResponse.CountPrognosaDpk = minidashboardResponse.CountPosisiDPK - minidashboardResponse.CountDpkJanjiStor - minidashboardResponse.CountDpkObAgf - +minidashboardResponse.CountDpkTidakBayar + minidashboardResponse.CountDpkBelumJatuhTempo

	return &response.BaseResponse{
		Success: true,
		Message: "successfully to list dpk",
		Data:    minidashboardResponse,
	}
}
