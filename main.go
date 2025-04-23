package main

import (
	"net/http"

	_ "go-swaggo-errorcontract/docs"
	"go-swaggo-errorcontract/handlers"
	"go-swaggo-errorcontract/handlers/middleware"
	"go-swaggo-errorcontract/models"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Error Contract API
// @version 1.16.4
// @description This is a sample server for error contract.
// @host localhost:8080
// @BasePath /
// @schemes http

var log = logrus.New()

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	log.Error("HTTP Error: ", err)
	c.JSON(code, models.ErrorContract{
		Code:    code,
		UserMsg: "An error occurred. Please try again later.",
		SysMsg:  err.Error(),
		Time:    c.Request().Header.Get(echo.HeaderXRequestID),
		DocsURL: "https://example.com/docs/errors",
	})
}

func main() {
	// Setup logger
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	})
	log.SetLevel(logrus.DebugLevel)

	// Initialize Echo
	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Use(mw.Recover())
	e.Use(mw.LoggerWithConfig(mw.LoggerConfig{
		Output: log.Writer(),
	}))
	e.Use(middleware.AuthenticationMiddleware())

	// Routes
	e.GET("/", func(c echo.Context) error {
		log.Debug("[Debug] Root endpoint accessed")
		log.Info("[Info] Root endpoint accessed")
		log.Warn("[Warn] Root endpoint accessed")
		log.Error("[Error] Root endpoint accessed")
		return c.String(http.StatusOK, "Welcome to the Error Contract API!")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/users", handlers.GetUsers)
	e.POST("/users", handlers.CreateUser)
	e.POST("/submit-form", handlers.SubmitForm)

	e.GET("/error", func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	})

	// Start server
	log.Info("Starting server on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
