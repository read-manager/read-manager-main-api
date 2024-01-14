package health

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

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
