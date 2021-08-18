package channels

import (
	"jomblo/business/channels"
	controller "jomblo/controllers"
	"jomblo/controllers/channels/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ChannelsController struct {
	ChannelsUsecase channels.Usecase
}

func NewChannelsController(Cc channels.Usecase) *ChannelsController {
	return &ChannelsController{
		ChannelsUsecase: Cc,
	}
}

func (ctrl *ChannelsController) ChannelsUserID(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Channels{}

	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	_, err := ctrl.ChannelsUsecase.ChannelsUserID(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, "success")
}
