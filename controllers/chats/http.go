package chats

import (
	"jomblo/business/chats"
	controller "jomblo/controllers"
	"jomblo/controllers/chats/request"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type ChatsController struct {
	ChatsUsecase chats.Usecase
}

func NewUserController(cc chats.Usecase) *ChatsController {
	return &ChatsController{
		ChatsUsecase: cc,
	}
}

func (ctrl *ChatsController) ChannelChatsID(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Chats{}

	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	_, err := ctrl.ChatsUsecase.ChannelChatsID(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, "success")
}
