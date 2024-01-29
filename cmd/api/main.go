package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
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
    db, dbErr := openDB(os.Getenv("POSTGRES_DSN"))
    if dbErr != nil {
        e.Logger.Fatal(dbErr.Error())
        os.Exit(1)
    }
    defer db.Close()
	infra.SetRoutes(e)
    go func() {
		if err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
    maxOpenConns, maxOpenConnsErr := strconv.Atoi(os.Getenv("POSTGRES_DB_MAX_OPEN_CONNS"))
	if maxOpenConnsErr != nil {
		os.Exit(1)
	}
	maxIdleConns, maxIdleConnsErr := strconv.Atoi(os.Getenv("POSTGRES_DB_MAX_IDLE_CONNS"))
	if maxIdleConnsErr != nil {
		os.Exit(1)
	}
	maxIdleTime, maxIdleTimeErr := time.ParseDuration(os.Getenv("POSTGRES_DB_MAX_IDLE_TIME"))
	if maxIdleTimeErr != nil {
		os.Exit(1)
	}
    db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(maxIdleTime)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
