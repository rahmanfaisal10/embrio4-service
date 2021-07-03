package handler

import (
	"rahmanfaisal10/embrio4-service/pkg/service"
	"rahmanfaisal10/embrio4-service/pkg/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRouter(s service.Service) *echo.Echo {
	// cfg := config.Get()

	// middleWareConfig := middleware.JWTConfig{
	// 	SigningKey: []byte(cfg.SecretKey),
	// 	Claims:     &request.JwtCustomClaims{},
	// }

	router := echo.New()
	router.Use(middleware.CORS(), middleware.Logger())

	//validator request
	router.Validator = &util.Validator{Validator: validator.New()}

	//router without authorization and authentication
	router.POST("api/v1/login", loginHandler(s))
	router.POST("api/v1/register", registerHandler(s))

	// authorizedRouter := router.Group("/api/v1/")
	// authorizedRouter.Use(middleware.JWTWithConfig(middleWareConfig))

	return router
}
