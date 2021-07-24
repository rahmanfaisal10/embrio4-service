package handler

import (
	"net/http"
	"rahmanfaisal10/embrio4-service/pkg/request"
	"rahmanfaisal10/embrio4-service/pkg/response"
	"rahmanfaisal10/embrio4-service/pkg/service"

	"github.com/labstack/echo"
)

func changePasswordHandler(svc service.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		resp := new(response.BaseResponse)
		req := new(request.ChangePasswordRequest)

		if err := c.Bind(req); err != nil {
			resp.Success = false
			resp.Message = err.Error()
			return c.JSON(http.StatusBadRequest, resp)
		}

		//validate request
		if err := c.Validate(req); err != nil {
			resp.Success = false
			resp.Message = err.Error()
			return c.JSON(http.StatusBadRequest, resp)
		}

		resp = svc.ChangedPasswordService(*req)
		return c.JSON(http.StatusOK, resp)
	}
}
