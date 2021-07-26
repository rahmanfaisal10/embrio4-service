package service

import (
	"rahmanfaisal10/embrio4-service/pkg/request"
	"rahmanfaisal10/embrio4-service/pkg/response"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/gommon/log"
)

const (
	SHEETNAME = "Sheet1"
)

var (
	uploadFiles = make([]request.FileRequest, 0)
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

	//get all value in row
	for _, row := range xlsx.GetRows(firstSheet) {
		getData := &request.FileRequest{
			Periode:                    row[0],
			Branch:                     row[1],
			Currency:                   row[2],
			NamaAO:                     row[3],
			LNType:                     row[4],
			NomorRekening:              row[5],
			NamaDebitur:                row[6],
			Plafond:                    row[7],
			NextPmtDate:                row[8],
			NextIntPmtDate:             row[9],
			Rate:                       row[10],
			TglMenunggak:               row[11],
			TglRealisasi:               row[12],
			TglJatuhTempo:              row[13],
			JangkaWaktu:                row[14],
			FlagRestruk:                row[15],
			CIFNO:                      row[16],
			KolektibilitasLancar:       row[17],
			KolektibilitasDPK:          row[18],
			KolektibilitasKurangLancar: row[19],
			KolektibilitasDiragukan:    row[20],
			KolektibilitasMacet:        row[21],
			TunggakanPokok:             row[22],
			TunggakanBunga:             row[23],
			TunggakanPinalty:           row[24],
			PN:                         row[25],
			NamaPN:                     row[26],
			Code:                       row[27],
			Description:                row[28],
			Kol_ADK:                    row[29],
			AvgOSHarian:                row[30],
			KecamatanTempatTinggal:     row[31],
			KelurahanTempatTinggal:     row[32],
			KodePosTempatTinggal:       row[33],
			KecamatanTempatUsaha:       row[34],
			KelurahanTempatUsaha:       row[35],
			KodePosTempatUsaha:         row[36],
		}

		//input in array object
		uploadFiles = append(uploadFiles, *getData)
	}

	return &response.BaseResponse{
		Success: true,
		Message: "successfully to upload file",
	}
}
