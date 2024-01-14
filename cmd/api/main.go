package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	infra "github.com/read-manager/read-manager-main-api/internal/infra/http"
)

// @title Read Manager Main API
// @version 1.0
// @description This is a API for a reader manager.
// @contact.name API Support
// @contact.email gustavocs789@gmail.com
// @host https://read-manager-main-api.onrender.com
// @BasePath /v1
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
    e.Use(middleware.CORS())
	if os.Getenv("LOAD_ENV_FILE") == "true" {
		loadEnvErr := godotenv.Load()
		if loadEnvErr != nil {
			e.Logger.Fatal(loadEnvErr.Error())
			os.Exit(1)
		}
	}
	infra.SetRoutes(e)
    go func() {
		if err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
