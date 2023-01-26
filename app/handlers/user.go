package handlers

import (
	"github.com/Serbroda/distance-challenge/security"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
}

func RegisterUserHandlers(e *echo.Group, h *UserHandler, baseUrl string) {
	e.GET(baseUrl+"/me", h.Me)
}

func (h *UserHandler) Me(ctx echo.Context) error {
	u, err := h.getUser(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, u)
}

func (si *UserHandler) getUser(ctx echo.Context) (string, error) {
	obj := ctx.Get("user")
	token := obj.(*jwt.Token)
	claims := token.Claims.(*security.Claims)
	if claims == nil {
		return "", echo.ErrUnauthorized
	}
	return claims.Subject, nil
}
