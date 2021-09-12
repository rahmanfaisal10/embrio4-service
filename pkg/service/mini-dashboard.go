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
		minidashboardResponse.Periode = v.Periode
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

	minidashboardResponse.PercenDpkBelumJatuhTempo = percen(float64(getAllJatuhTempo.Count), float64(getAllJatuhTempo.CountTotal))
	minidashboardResponse.PercenDpkJanjiStor = percen(minidashboardResponse.DpkJanjiStor, minidashboardResponse.PosisiDPK)
	minidashboardResponse.PercenDpkObAgf = percen(minidashboardResponse.DpkObAgf, minidashboardResponse.PosisiDPK)
	minidashboardResponse.PercenDpkTidakBayar = percen(minidashboardResponse.DpkTidakBayar, minidashboardResponse.PosisiDPK)
	minidashboardResponse.PercenPosisiDPK = percen(minidashboardResponse.PosisiDPK, minidashboardResponse.PosisiDPK)
	minidashboardResponse.PercenPrognosaDpk = percen(float64(minidashboardResponse.CountPrognosaDpk), getAllJatuhTempo.DpkTotal)

	return &response.BaseResponse{
		Success: true,
		Message: "successfully to list dpk",
		Data:    minidashboardResponse,
	}
}

func percen(pembilang, penyebut float64) float64 {
	return pembilang / penyebut * 100
}
