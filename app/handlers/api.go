package handlers

import (
	"github.com/Serbroda/distance-challenge/models"
	"github.com/Serbroda/distance-challenge/security"
	"github.com/Serbroda/distance-challenge/services"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserChange struct {
	Password string `json:"password" form:"password"`
}

type UserHandler struct {
	UserService *services.UserService
}

func RegisterUserHandlers(e *echo.Group, h *UserHandler, baseUrl string) {
	e.GET(baseUrl+"/me", h.Me)
	e.GET(baseUrl+"/users", h.GetUsers)
	e.GET(baseUrl+"/users/:id", h.GetUser)
	e.PUT(baseUrl+"/users/:id", h.UpdateUser)
	e.GET(baseUrl+"/users/:id/activate", h.Activate)
}

func (h *UserHandler) Me(ctx echo.Context) error {
	u, err := h.getUser(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, u)
}

func (h *UserHandler) GetUsers(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, h.UserService.GetAll())
}

func (h *UserHandler) GetUser(ctx echo.Context) error {
	userId64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return echo.ErrInternalServerError
	}
	userId := uint(userId64)

	user, err := h.UserService.Get(userId)
	if err != nil {
		return echo.ErrNotFound
	}
	return ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	u, err := h.getUser(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}

	userId, err := h.parseParamUint(ctx, "id")
	if err != nil {
		return err
	}

	var payload UserChange
	err = ctx.Bind(&payload)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	if u.ID != userId {
		return echo.ErrForbidden
	}

	user, err := h.UserService.Get(userId)
	if err != nil {
		return echo.ErrNotFound
	}
	user.Password = security.MustHashPassword(payload.Password)
	user, err = h.UserService.Update(user)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, user)
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

func (si *UserHandler) getUser(ctx echo.Context) (models.User, error) {
	obj := ctx.Get("user")
	token := obj.(*jwt.Token)
	claims := token.Claims.(*security.Claims)
	if claims == nil {
		return models.User{}, echo.ErrUnauthorized
	}
	return si.UserService.Get(claims.UserId)
}

func (si *UserHandler) parseParamUint(ctx echo.Context, param string) (uint, error) {
	userId64, err := strconv.ParseUint(ctx.Param(param), 10, 64)
	if err != nil {
		return -1, echo.ErrInternalServerError
	}
	return uint(userId64), nil
}
