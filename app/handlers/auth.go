package handlers

import (
	"github.com/Serbroda/distance-challenge/models"
	"github.com/Serbroda/distance-challenge/security"
	"github.com/Serbroda/distance-challenge/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	UserService *services.UserService
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RegistrationRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func RegisterAuthHandlers(e *echo.Echo, h *AuthHandler, baseUrl string) {
	e.POST(baseUrl+"/auth/login", h.Login)
	e.POST(baseUrl+"/auth/register", h.Register)
}

// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "login body"
// @Success 200 {object} TokenPair
// @Failure 400
// @Router /auth/login [post]
func (h *AuthHandler) Login(ctx echo.Context) error {
	var payload LoginRequest
	err := ctx.Bind(&payload)
	if err != nil || payload.Username == "" || payload.Password == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	u, err := h.UserService.GetByUsername(payload.Username)

	if err != nil || !security.CheckPasswordHash(payload.Password, u.Password) || !u.Active {
		return ctx.String(http.StatusBadRequest, "invalid login")
	}

	tokens, err := security.GenerateJwtPair("s3cret", u.Username, u.ID, 10, 600)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to generate token")
	}

	return ctx.JSON(http.StatusOK, tokens)
}

func (h *AuthHandler) Register(ctx echo.Context) error {
	var payload RegistrationRequest
	err := ctx.Bind(&payload)
	if err != nil || payload.Username == "" || payload.Password == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	_, err = h.UserService.GetByUsername(payload.Username)

	if err == nil {
		return ctx.String(http.StatusBadRequest, "user already exists")
	}

	user, err := h.UserService.Create(models.User{
		Username: payload.Username,
		Password: security.MustHashPassword(payload.Password),
		Active:   false,
	})

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed creating user")
	}

	return ctx.JSON(http.StatusCreated, user)
}
