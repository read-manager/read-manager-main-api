package health

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// CheckHandler godoc
// @Summary      Health check
// @Description  Check health of API
// @Tags         health
// @Produce      json
// @Success      200 {object} checResponse
// @Router       /v1/healthcheck [get]
func CheckHandler(c echo.Context) error {
    data := map[string]any{
        "status": "available",
		"system_info": map[string]string{
			"environment": os.Getenv("ENV"),
			"version":     os.Getenv("VERSION"),
		},
    }
    return c.JSON(http.StatusOK, data)
}
