package service

import (
	"rahmanfaisal10/embrio4-service/pkg/model"
	"rahmanfaisal10/embrio4-service/pkg/request"
	"rahmanfaisal10/embrio4-service/pkg/response"
	"rahmanfaisal10/embrio4-service/pkg/util"
	"time"

	"github.com/labstack/gommon/log"
)

func (s *service) InsertLogTandaiService(request *request.LogTandaiRequest) *response.BaseResponse {
	//periode
	periode, err := util.ParseStringToDate("2006-01-02", request.Periode)
	if err != nil {
		log.Error(err)
		return &response.BaseResponse{
			Success: false,
			Message: "failed to parse field periode from string to date",
		}
	}

	//periode
	tglJanjiSetor, err := util.ParseStringToDate("2006-01-02", request.TglJanjiSetor)
	if err != nil {
		log.Error(err)
		return &response.BaseResponse{
			Success: false,
			Message: "failed to parse field periode from string to date",
		}
	}

	logTandai := &model.LogTandai{
		PnPengguna:            request.PnPengguna,
		NamaPenguna:           request.NamaPenguna,
		Periode:               periode,
		CifNo:                 request.CifNo,
		NomorRekeningSimpanan: request.NomorRekeningSimpanan,
		NomorRekeningPinjaman: request.NomorRekeningPinjaman,
		Status:                request.Status,
		TglJanjiSetor:         tglJanjiSetor,
		CreatedAt:             &time.Time{},
		UpdatedAt:             &time.Time{},
	}

	err = s.r.InsertLogTandai(logTandai)
	if err != nil {
		log.Error(err)
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	return &response.BaseResponse{
		Success: true,
		Message: "successfully to insert log tandai",
	}
}
