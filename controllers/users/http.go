package users

import (
	"jomblo/business/users"
	controller "jomblo/controllers"
	"jomblo/controllers/users/request"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	UserUsecase users.Usecase
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{
		UserUsecase: uc,
	}
}

func (ctrl *UserController) Regis(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Users{}

	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	_, err := ctrl.UserUsecase.Regis(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, "success")
}
