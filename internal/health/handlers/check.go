package health

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type System_info struct {
	Environment string `json:"environment"`
	Version string `json:"version"`
}

type CheckResponse struct {
	Status string `json:"status"`
	System_info System_info `json:"system_info"`
}

// CheckHandler godoc
// @Summary      Health check
// @Description  Check health of API
// @Tags         health
// @Produce      json
// @Success      200 {object} CheckResponse
// @Router       /v1/healthcheck [get]
func CheckHandler(c echo.Context) error {
    data := CheckResponse{
        Status: "available",
		System_info: System_info{
			Environment: os.Getenv("ENV"),
			Version: os.Getenv("VERSION"),
		},
    }
    return c.JSON(http.StatusOK, data)
}
