package handlers

import (
	"net/http"

	"go-swaggo-errorcontract/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	_ "github.com/swaggo/echo-swagger"
)

var users = []models.User{
	{Name: "John Doe", Email: "john@mail.com"},
	{Name: "Jane Doe", Email: "jane@mail.com"},
}

// GetUsers godoc
// @Summary Get users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} models.SuccessResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /users [get]
func GetUsers(c echo.Context) error {
	if len(users) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, models.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "No users found",
		})
	}

	log.Info("Users retrieved successfully")
	return c.JSON(http.StatusOK, models.SuccessResponse{
		Data: users,
	})
}

// CreateUser godoc
// @Summary Create user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users [post]
func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid input",
		})
	}

	if user.Name == "" || user.Email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Name and Email are required",
		})
	}

	users = append(users, user)
	log.Info("User created successfully")
	return c.JSON(http.StatusCreated, models.SuccessResponse{
		Data: user,
	})
}

// SubmitForm godoc
// @Summary Submit form
// @Description Submit a form
// @Tags form
// @Accept json
// @Produce json
// @Param name formData string true "Name"
// @Param email formData string true "Email"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /form [post]
func SubmitForm(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	if name == "" || email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Name and Email are required",
		})
	}

	log.Info("Form submitted successfully")
	return c.JSON(http.StatusOK, models.SuccessResponse{
		Data: map[string]string{
			"name":  name,
			"email": email,
		},
	})
}
