package handler

import (
	"net/http"
	"rahmanfaisal10/embrio4-service/pkg/service"

	"github.com/labstack/echo"
)

func targetInsertHandler(s service.Service) func(c echo.Context) error {
	return func(c echo.Context) error {

		token := c.Request().Header.Get("Authorization")

		resp := s.InsertAutoTarget(token)
		if !resp.Success {
			return c.JSON(http.StatusBadRequest, resp)
		}
		return c.JSON(http.StatusOK, resp)
	}
}
