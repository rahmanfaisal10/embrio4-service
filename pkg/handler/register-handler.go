package handler

import (
	"net/http"
	"rahmanfaisal10/embrio4-service/pkg/response"
	"rahmanfaisal10/embrio4-service/pkg/service"

	"github.com/labstack/echo"
)

func registerHandler(s service.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		resp := new(response.BaseResponse)

		token := c.Request().Header.Get("Authorization")

		err := s.RegisterService(token)
		if err != nil {
			resp.Success = false
			resp.Message = err.Error()
			return c.JSON(http.StatusBadRequest, resp)
		}

		resp.Success = true
		resp.Message = "Register Success"
		return c.JSON(http.StatusOK, resp)
	}
}
