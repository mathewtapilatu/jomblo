package matches

import (
	"jomblo/business/matches"
	controller "jomblo/controllers"
	"jomblo/controllers/matches/request"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type MatchesController struct {
	MatchesUsecase matches.Usecase
}

func NewMatchesController(Mc matches.Usecase) *MatchesController {
	return &MatchesController{
		MatchesUsecase: Mc,
	}
}

// func (ctrl *MatchesController) UserID(c echo.Context) error {
// 	ctx := c.Request().Context()
// 	userID, _ := strconv.Atoi(c.Param("id"))

// 	_, err := ctrl.MatchesUsecase.UserID(ctx, userID)
// 	if err != nil {
// 		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}
// 	return controller.NewSuccessResponse(c, "success")
// }

func (ctrl *MatchesController) UserID(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Matches{}

	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	_, err := ctrl.MatchesUsecase.UserID(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, "success")
}
