package service

import (
	"rahmanfaisal10/embrio4-service/pkg/model"
	"rahmanfaisal10/embrio4-service/pkg/response"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/gommon/log"
)

const (
	SHEETNAME = "Sheet1"
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
		getData := &model.Upload{
			Periode:                    row[0],
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
			NextPmtDate:                row[14],
			NextIntPmtDate:             row[15],
			Rate:                       row[16],
			TglMenunggak:               row[17],
			TglRealisasi:               row[18],
			TglJatuhTempo:              row[19],
			JangkaWaktu:                row[20],
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

	return &response.BaseResponse{
		Success: true,
		Message: "successfully to upload file",
	}
}
