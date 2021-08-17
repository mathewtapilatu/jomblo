package routes

import (
	"jomblo/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Regis)
}

func (cl *ControllerList) RouteMatches(e *echo.Echo) {
	users := e.Group("matches")
	users.POST("/matches", cl.UserController.Regis)
}
