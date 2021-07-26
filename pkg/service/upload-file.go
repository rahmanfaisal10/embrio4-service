package service

import (
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/response"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func (svc *service) ImportFileUploadDWH(destination string) *response.BaseResponse {
	xlsx, err := excelize.OpenFile(destination)
	if err != nil {
		return &response.BaseResponse{
			Success: false,
			Message: "sdasdsa",
		}
	}

	fmt.Println(xlsx.GetCellValue("SheetOne", "A4"))

	return &response.BaseResponse{
		Success: true,
		Message: "successfully to upload file",
	}
}
