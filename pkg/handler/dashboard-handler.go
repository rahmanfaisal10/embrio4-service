package handler

import (
	"net/http"
	"rahmanfaisal10/embrio4-service/pkg/service"

	"github.com/labstack/echo"
)

func dashboardHandler(s service.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		mantriRequest := c.Param("mantri")

		resp := s.ViewDashboardService(mantriRequest)

		if !resp.Success {
			return c.JSON(http.StatusBadRequest, resp)
		}
		return c.JSON(http.StatusOK, resp)
	}
}
