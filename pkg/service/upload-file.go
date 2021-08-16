package service

import (
	"fmt"
	"math"
	"rahmanfaisal10/embrio4-service/pkg/model"
	"rahmanfaisal10/embrio4-service/pkg/response"
	"rahmanfaisal10/embrio4-service/pkg/util"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/gommon/log"
)

const (
	SHEETNAME    = "Sheet1"
	layoutISO    = "2006-01-02"
	formatUpload = "02/01/2006"
)

var (
	uploadModel = make([]*model.Upload, 0)
)

func (svc *service) ImportFileUploadDWH(destination string) *response.BaseResponse {
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
		//when get title, so loop to next rows
		if key == 0 {
			continue
		}

		//periode
		periode, err := util.ParseStringToDate(layoutISO, row[0])
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: "failed to parse field periode from string to date",
			}
		}

		//next payment date
		nextPmtDate, err := util.ParseStringToDate(formatUpload, row[8])
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: "failed to parse field next_pmt_date from string to date",
			}
		}

		//next int pmt date
		nextIntPmtDate, err := util.ParseStringToDate(formatUpload, row[9])
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: "failed to parse field next_int_pmt_date from string to date",
			}
		}

		//tgl menunggak
		tglMenunggak, err := util.ParseStringToDate(formatUpload, row[11])
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: "failed to parse field tgl_menunggak from string to date",
			}
		}

		//next realisasi
		TglRealisasi, err := util.ParseStringToDate(formatUpload, row[12])
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: "failed to parse field tgl_realisasi from string to date",
			}
		}

		//jatuh tempo
		tglJatuhTempo, err := util.ParseStringToDate(formatUpload, row[13])
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: "failed to parse field tgl_jatuh_tempo from string to date",
			}
		}

		getData := &model.Upload{
			Periode:                    periode,
			Branch:                     row[1],
			Currency:                   row[2],
			NamaAO:                     row[3],
			LNType:                     row[4],
			NomorRekening:              row[5],
			NamaDebitur:                row[6],
			Plafond:                    util.ParseStringToFloat(row[7]),
			NextPmtDate:                nextPmtDate,
			NextIntPmtDate:             nextIntPmtDate,
			Rate:                       float32(util.ParseStringToFloat(row[10])),
			TglMenunggak:               tglMenunggak,
			TglRealisasi:               TglRealisasi,
			TglJatuhTempo:              tglJatuhTempo,
			JangkaWaktu:                row[14],
			FlagRestruk:                row[15],
			CIFNO:                      row[16],
			KolektibilitasLancar:       util.ParseStringToFloat(row[17]),
			KolektibilitasDPK:          util.ParseStringToFloat(row[18]),
			KolektibilitasKurangLancar: util.ParseStringToFloat(row[19]),
			KolektibilitasDiragukan:    util.ParseStringToFloat(row[20]),
			KolektibilitasMacet:        util.ParseStringToFloat(row[21]),
			TunggakanPokok:             util.ParseStringToFloat(row[22]),
			TunggakanBunga:             util.ParseStringToFloat(row[23]),
			TunggakanPinalty:           util.ParseStringToFloat(row[24]),
			PNPengelola:                row[25],
			NamaPengelola:              row[26],
			Code:                       row[27],
			Description:                row[28],
			KolADK:                     int64(util.ParseStringToFloat(row[29])),
			AvgOsHarian:                row[30],
			KecamatanTempatTinggal:     row[31],
			KelurahanTempatTinggal:     row[32],
			KodePosTempatTinggal:       row[33],
			KecamatanTempatUsaha:       row[34],
			KelurahanTempatUsaha:       row[35],
			KodePosTempatUsaha:         row[36],
		}

		uploadModel = append(uploadModel, getData)
	}

	err = svc.r.UploadRepository(uploadModel)
	if err != nil {
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	//list mantri
	err = svc.r.InsertDashboard()
	if err != nil {
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	duration := time.Since(time.Now())
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")

	return &response.BaseResponse{
		Success: true,
		Message: "successfully to upload file",
	}
}
