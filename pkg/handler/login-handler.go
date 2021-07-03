package handler

import (
	"net/http"
	"rahmanfaisal10/embrio4-service/pkg/request"
	"rahmanfaisal10/embrio4-service/pkg/response"
	"rahmanfaisal10/embrio4-service/pkg/service"

	"github.com/labstack/echo"
)

func loginHandler(s service.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		resp := new(response.BaseResponse)
		req := &request.LoginRequest{
			UsernamePN: c.FormValue("username_pn"),
			Password:   c.FormValue("password"),
		}

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

		//create token and get data from service
		createToken, err := s.LoginService(*req)
		if err != nil {
			resp.Success = false
			resp.Message = err.Error()
			return c.JSON(http.StatusUnauthorized, resp)
		}

		resp.Success = true
		resp.Message = "Success login"
		resp.Data = createToken
		return c.JSON(http.StatusOK, resp)
	}
}
