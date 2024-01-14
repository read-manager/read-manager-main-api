package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/read-manager/read-manager-main-api/internal/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	if os.Getenv("LOAD_ENV_FILE") == "true" {
		loadEnvErr := godotenv.Load()
		if loadEnvErr != nil {
			e.Logger.Fatal(loadEnvErr.Error())
			os.Exit(1)
		}
	}
	http.SetRoutes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
