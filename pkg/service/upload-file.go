package service

import (
	"fmt"
	"math"
	"rahmanfaisal10/embrio4-service/pkg/model"
	"rahmanfaisal10/embrio4-service/pkg/response"
	"strconv"
	"strings"
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
		dateParse, err := time.Parse(formatUpload, row[0])
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: err.Error(),
			}
		}
		oriDate := dateParse.Format("2006-01-02")
		periode, err := time.Parse("2006-01-02", oriDate)
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: err.Error(),
			}
		}

		//next payment date
		dateParse, err = time.Parse(formatUpload, row[14])
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: err.Error(),
			}
		}
		oriDate = dateParse.Format("2006-01-02")
		nextPmtDate, err := time.Parse("2006-01-02", oriDate)
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: err.Error(),
			}
		}

		//next realisasi
		dateParse, err = time.Parse(formatUpload, row[18])
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: err.Error(),
			}
		}
		oriDate = dateParse.Format("2006-01-02")
		TglRealisasi, err := time.Parse("2006-01-02", oriDate)
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: err.Error(),
			}
		}

		//jatuh tempo
		dateParse, err = time.Parse(formatUpload, row[19])

		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: err.Error(),
			}
		}
		oriDate = dateParse.Format("2006-01-02")
		TglJatuhTempo, err := time.Parse("2006-01-02", oriDate)
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: err.Error(),
			}
		}

		//jangkaWaktu
		stringJangkaWaktu := strings.ReplaceAll(row[20], "M", "")
		jangkaWaktu, err := strconv.Atoi(stringJangkaWaktu)
		if err != nil {
			log.Error(err)
			return &response.BaseResponse{
				Success: false,
				Message: err.Error(),
			}
		}

		getData := &model.Upload{
			Periode:                    periode,
			Region:                     row[1],
			Mainbranch:                 row[2],
			Branch:                     row[3],
			Currency:                   row[4],
			NamaAO:                     row[5],
			LNType:                     row[6],
			NomorRekening:              row[7],
			NamaDebitur:                row[8],
			AlamatIdentitas:            row[9],
			KodePosIdentitas:           row[10],
			AlamatKantor:               row[11],
			KodePosKantor:              row[12],
			Plafond:                    row[13],
			NextPmtDate:                nextPmtDate,
			NextIntPmtDate:             row[15],
			Rate:                       row[16],
			TglMenunggak:               row[17],
			TglRealisasi:               TglRealisasi,
			TglJatuhTempo:              TglJatuhTempo,
			JangkaWaktu:                jangkaWaktu,
			FlagRestruk:                row[21],
			CIFNO:                      row[22],
			KolektibilitasLancar:       row[23],
			KolektibilitasDPK:          row[24],
			KolektibilitasKurangLancar: row[25],
			KolektibilitasDiragukan:    row[26],
			KolektibilitasMacet:        row[27],
			TunggakanPokok:             row[28],
			TunggakanBunga:             row[29],
			TunggakanPinalty:           row[30],
			PNPengelola:                row[31],
			NamaPengelola:              row[32],
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
	_, err = svc.r.GetAllMantri()
	if err != nil {
		log.Error(err)
		return &response.BaseResponse{
			Success: false,
			Message: err.Error(),
		}
	}

	// for _, v := range mantri {
	// 	//execute query total os
	// 	_, err := svc.r.QueryTotalOS(v.Kode)
	// 	if err != nil {
	// 		log.Error(err)
	// 		return &response.BaseResponse{
	// 			Success: false,
	// 			Message: err.Error(),
	// 		}
	// 	}

	// 	//sisaSuplesi
	// 	_, err = svc.r.QuerySisaSuplesi(v.Kode)
	// 	if sql.ErrNoRows == err {
	// 		err = nil
	// 	}

	// 	if err != nil {
	// 		log.Error(err)
	// 		return &response.BaseResponse{
	// 			Success: false,
	// 			Message: err.Error(),
	// 		}
	// 	}

	// 	//sisa lunas hutang
	// 	_, err = svc.r.QueryRincianLunasHutang(v.Kode)
	// 	if err != nil {
	// 		log.Error(err)
	// 		return &response.BaseResponse{
	// 			Success: false,
	// 			Message: err.Error(),
	// 		}
	// 	}
	// }

	duration := time.Since(time.Now())
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")

	return &response.BaseResponse{
		Success: true,
		Message: "successfully to upload file",
	}
}
