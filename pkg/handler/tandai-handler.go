package handler

import (
	"net/http"
	"rahmanfaisal10/embrio4-service/pkg/request"
	"rahmanfaisal10/embrio4-service/pkg/response"
	"rahmanfaisal10/embrio4-service/pkg/service"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func insertTandaiHandler(s service.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		resp := new(response.BaseResponse)
		req := new(request.LogTandaiRequest)

		if err := c.Bind(req); err != nil {
			log.Error(err)
			resp.Success = false
			resp.Message = err.Error()
			return c.JSON(http.StatusBadRequest, resp)
		}

		//validate request
		if err := c.Validate(req); err != nil {
			log.Error(err)
			resp.Success = false
			resp.Message = err.Error()
			return c.JSON(http.StatusBadRequest, resp)
		}

		resp = s.InsertLogTandaiService(req)
		if !resp.Success {
			return c.JSON(http.StatusBadRequest, resp)
		}
		return c.JSON(http.StatusOK, resp)
	}
}
