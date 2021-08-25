package handler

import (
	"net/http"
	"rahmanfaisal10/embrio4-service/config"
	"rahmanfaisal10/embrio4-service/pkg/request"
	"rahmanfaisal10/embrio4-service/pkg/service"
	"rahmanfaisal10/embrio4-service/pkg/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRouter(s service.Service) *echo.Echo {
	cfg := config.Get()

	middleWareConfig := middleware.JWTConfig{
		SigningKey: []byte(cfg.SecretKey),
		Claims:     &request.JwtCustomClaims{},
	}

	router := echo.New()
	router.Use(middleware.CORS(), middleware.Logger())

	//validator request
	router.Validator = &util.Validator{Validator: validator.New()}

	//router without authorization and authentication
	router.POST("api/login", loginHandler(s))
	router.GET("api", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Embrio4 server connected")
	})

	authorizedRouter := router.Group("/api/v1/")
	authorizedRouter.Use(middleware.JWTWithConfig(middleWareConfig))
	authorizedRouter.POST("email-verification", emailVerificationHandler(s))
	authorizedRouter.POST("register", registerHandler(s))
	authorizedRouter.PUT("change-password", changePasswordHandler(s))
	authorizedRouter.POST("upload-file", uploadFileHandler(s))
	authorizedRouter.GET("dashboard/:mantri", dashboardHandler(s))
	authorizedRouter.POST("upload-file-simpanan", uploadFileSimpananHandler(s))
	authorizedRouter.POST("tandai-activity", insertTandaiHandler(s))
	authorizedRouter.GET("list-dpk/:mantri", listDpkHandler(s))
	authorizedRouter.GET("minidashboard/:mantri", miniDashboardHandler(s))

	return router
}
