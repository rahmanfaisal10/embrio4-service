package handler

import (
	"net/http"
	"rahmanfaisal10/embrio4-service/pkg/service"

	"github.com/labstack/echo"
)

func listDpkHandler(service service.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		mantriRequest := c.Param("mantri")

		resp := service.ListDpkService(mantriRequest)
		if !resp.Success {
			return c.JSON(http.StatusBadRequest, resp)
		}

		return c.JSON(http.StatusOK, resp)
	}
}
