package handlers

import (
	"github.com/Serbroda/distance-challenge/models"
	"github.com/Serbroda/distance-challenge/security"
	"github.com/Serbroda/distance-challenge/services"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	UserService *services.UserService
}

func RegisterUserHandlers(e *echo.Group, h *UserHandler, baseUrl string) {
	e.GET(baseUrl+"/me", h.Me)
	e.GET(baseUrl+"/users/:id/activate", h.Activate)
}

func (h *UserHandler) Activate(ctx echo.Context) error {
	u, err := h.getUser(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}
	if !u.IsAdmin {
		return echo.NewHTTPError(http.StatusForbidden, "no admin")
	}
	userId := ctx.Param("id")
	return ctx.String(http.StatusOK, "user "+userId+" activated")
}

func (h *UserHandler) Me(ctx echo.Context) error {
	u, err := h.getUser(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, u)
}

func (si *UserHandler) getUser(ctx echo.Context) (models.User, error) {
	obj := ctx.Get("user")
	token := obj.(*jwt.Token)
	claims := token.Claims.(*security.Claims)
	if claims == nil {
		return models.User{}, echo.ErrUnauthorized
	}
	return si.UserService.Get(claims.UserId)
}
