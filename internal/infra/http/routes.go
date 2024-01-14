package infra

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"

	health "github.com/read-manager/read-manager-main-api/internal/health/handlers"
)

func SetRoutes(e *echo.Echo) {
	e.GET("/v1/docs/*", echoSwagger.WrapHandler)
	e.GET("/v1/healthcheck", health.CheckHandler)
}
