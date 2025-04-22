package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "go-swaggo-errorcontract/docs" // Swagger generated files
)

// @title Echo Swagger Full Example API
// @version 1.0
// @description Sample API with JSON, form-data, Swagger, Logrus, and centralized error handling
// @host localhost:8080
// @BasePath /

var log = logrus.New()

// User represents a user object
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserForm represents a user submitted via form-data
type UserForm struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ErrorResponse standard error response
type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

var users = []User{
	{ID: 1, Name: "Alice"},
}

// getUsers godoc
// @Summary Get all users
// @Tags users
// @Produce json
// @Success 200 {array} User
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

// createUser godoc
// @Summary Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User to create"
// @Success 201 {object} User
// @Failure 400 {object} ErrorResponse
// @Router /users [post]
func createUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}
	u.ID = len(users) + 1
	users = append(users, *u)
	return c.JSON(http.StatusCreated, u)
}

// submitForm godoc
// @Summary Submit user via form-data
// @Tags form
// @Accept mpfd
// @Produce json
// @Param name formData string true "User name"
// @Param age formData int true "User age"
// @Success 200 {object} UserForm
// @Failure 400 {object} ErrorResponse
// @Router /submit-form [post]
func submitForm(c echo.Context) error {
	name := c.FormValue("name")
	ageStr := c.FormValue("age")

	if name == "" || ageStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Name and age are required")
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Age must be a number")
	}

	return c.JSON(http.StatusOK, UserForm{
		Name: name,
		Age:  age,
	})
}

// Custom error handler
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"

	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		code = httpErr.Code
		message = httpErr.Message.(string)
	}

	log.WithFields(logrus.Fields{
		"method": c.Request().Method,
		"path":   c.Path(),
		"error":  err.Error(),
	}).Error("Request failed")

	if !c.Response().Committed {
		c.JSON(code, ErrorResponse{
			Message: message,
			Code:    code,
		})
	}
}

func main() {
	// Setup logger
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(logrus.InfoLevel)

	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = customHTTPErrorHandler

	// Middleware
	e.Use(mw.LoggerWithConfig(mw.LoggerConfig{
		Output: log.Writer(),
	}))
	e.Use(mw.Recover())

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Routes
	e.GET("/users", getUsers)
	e.POST("/users", createUser)
	e.POST("/submit-form", submitForm)

	log.Info("🚀 Server running at http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
