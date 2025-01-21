package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"musiclib.com.co/musiclib/models"
	"musiclib.com.co/musiclib/services"
)

func GetAllUsers(c echo.Context) error {
	response, err := services.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, response)
}

func SaveUser(c echo.Context) error {
	u := new(models.User)
	fmt.Println(u)

	if err := c.Bind(u); err != nil {
		return err
	}

	defaultBaseTemplate := models.BaseTemplate{}

	baseTemplate := models.NewBaseTemplate(defaultBaseTemplate)
	u.BaseTemplate = *baseTemplate

	u, err := services.CreateUser(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, u)
}
