package service

import (
	"rahmanfaisal10/embrio4-service/pkg/model"
	"rahmanfaisal10/embrio4-service/pkg/response"
	"rahmanfaisal10/embrio4-service/pkg/util"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/gommon/log"
)

func (s *service) ImportFileSimpanan(destination string) *response.BaseResponse {
	simpanans := make([]*model.Simpanan, 0)

	xlsx, err := excelize.OpenFile(destination)
	if err != nil {
		log.Error(err)
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	firstSheet := xlsx.WorkBook.Sheets.Sheet[0].Name

	for key, row := range xlsx.GetRows(firstSheet) {
		if key == 0 {
			continue
		}

		//periode
		periode, err := util.ParseStringToDate(formatUpload, row[0])
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: "failed to parse field periode from string to date",
			}
		}

		//dlt
		dlt, err := util.ParseStringToDate("2/01/06", row[8])
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: "failed to parse field dlt from string to date",
			}
		}

		// opendt
		opendt, err := util.ParseStringToDate("1/2/2006", row[9])
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: "failed to parse field opendt from string to date",
			}
		}

		getData := &model.Simpanan{
			Periode:                periode,
			UkerCode:               row[1],
			CurrCode:               row[2],
			CUrrDesc:               row[3],
			AccountNumber:          row[4],
			CifNo:                  row[5],
			ShortName:              row[6],
			OfficerName:            row[7],
			Dlt:                    dlt,
			OpenDT:                 opendt,
			Balance:                util.ParseStringToFloat(row[10]),
			AvailableBalance:       util.ParseStringToFloat(row[11]),
			IntCredit:              util.ParseStringToFloat(row[12]),
			AccruedInt:             util.ParseStringToFloat(row[13]),
			AverageBalance:         util.ParseStringToFloat(row[14]),
			ProdCode:               row[15],
			PnPengelola:            int64(util.ParseStringToFloat(row[16])),
			NamaPengelola:          row[17],
			KecamatanTempatTinggal: row[18],
			KelurahanTempatTinggal: row[19],
			KodePosTempatTinggal:   row[20],
			KecamatanTempatUsaha:   row[21],
			KelurahanTempatUsaha:   row[22],
			KodePosTempatUsaha:     row[23],
		}

		simpanans = append(simpanans, getData)
	}

	err = s.r.InsertSimpanan(simpanans)
	if err != nil {
		log.Error(err)
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	return &response.BaseResponse{
		Success: true,
		Message: "successfully to insert file D319",
	}
}
