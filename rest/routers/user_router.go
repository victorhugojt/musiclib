package routers

import (
	"github.com/labstack/echo/v4"
	"musiclib.com.co/musiclib/rest/controllers"
)

var (
	router = echo.New()
)

func StartRouting() {
	router.GET("/users", controllers.GetAllUsers)
	router.POST("/users", controllers.SaveUser)
	router.POST("/libraries", controllers.SaveLibrary)
	router.GET("/libraries", controllers.GetAllLibraries)
	router.Logger.Fatal(router.Start(":3000"))
}
