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
	router.Logger.Fatal(router.Start(":3000"))
}
