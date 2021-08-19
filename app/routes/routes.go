package routes

import (
	"jomblo/controllers/channels"
	"jomblo/controllers/chats"
	"jomblo/controllers/matches"
	"jomblo/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController     users.UserController
	MatchesController  matches.MatchesController
	ChannelsController channels.ChannelsController
	ChatsController    chats.ChatsController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Regis)

	matches := e.Group("matches")
	matches.POST("/matches", cl.MatchesController.UserID)

	channels := e.Group("channels")
	channels.POST("/channels", cl.ChannelsController.ChannelsUserID)

	chats := e.Group("chats")
	chats.POST("/message", cl.ChatsController.ChannelChatsID)
}
