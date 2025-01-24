package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"musiclib.com.co/musiclib/models"
	"musiclib.com.co/musiclib/services"
)

func GetLibraryById(c echo.Context) error {
	response, err := services.GetLibraryById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, response)
}

func GetAllLibraries(c echo.Context) error {
	response, err := services.GetAllLibraries()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, response)
}

func SaveLibrary(c echo.Context) error {
	l := new(models.Library)
	fmt.Println(l)

	if err := c.Bind(l); err != nil {
		return err
	}

	defaultBaseTemplate := models.BaseTemplate{}

	baseTemplate := models.NewBaseTemplate(defaultBaseTemplate)
	l.BaseTemplate = *baseTemplate

	l, err := services.CreateLibrary(l)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, l)
}
