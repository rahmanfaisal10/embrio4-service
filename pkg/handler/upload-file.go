package handler

import (
	"io"
	"net/http"
	"os"
	"rahmanfaisal10/embrio4-service/pkg/response"
	"rahmanfaisal10/embrio4-service/pkg/service"

	"github.com/labstack/echo"
)

func uploadFileHandler(service service.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		resp := new(response.BaseResponse)

		//source
		file, err := c.FormFile("file")
		if err != nil {
			resp.Success = false
			resp.Message = err.Error()
			return c.JSON(http.StatusBadRequest, resp)
		}

		src, err := file.Open()
		if err != nil {
			resp.Success = false
			resp.Message = err.Error()
			return c.JSON(http.StatusBadRequest, resp)
		}

		// Destination
		dst, err := os.Create("pkg/assets/" + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		respo := service.ImportFileUploadDWH(dst.Name())
		if !respo.Success {
			respo.Success = false
			respo.Message = err.Error()
			return c.JSON(http.StatusBadRequest, resp)
		} else {
			resp.Success = respo.Success
			resp.Message = respo.Message
			return c.JSON(http.StatusOK, resp)
		}

	}
}
