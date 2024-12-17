package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"musiclib.com.co/musiclib/services"
)

func GetAllUsers(c echo.Context) error {
	response, err := services.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, response)
}
