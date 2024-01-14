package http

import (
	"github.com/labstack/echo/v4"
	health "github.com/read-manager/read-manager-main-api/internal/health/handlers"
)

func SetRoutes(e *echo.Echo) {
	e.GET("/v1/healthcheck", health.CheckHandler)
}
