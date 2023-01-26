package security

import (
	"github.com/Serbroda/distance-challenge/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Provider user.UserService
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RegistrationRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func RegisterHandlers(e *echo.Echo, h *Handler, baseUrl string) {
	e.POST(baseUrl+"/auth/signin", h.SignIn)
	e.POST(baseUrl+"/auth/signup", h.SignUp)
}

// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "login body"
// @Success 200 {object} TokenPair
// @Failure 400
// @Router /auth/login [post]
func (h *Handler) SignIn(ctx echo.Context) error {
	var payload LoginRequest
	err := ctx.Bind(&payload)
	if err != nil || payload.Username == "" || payload.Password == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	user, err := h.Provider.GetByUsername(payload.Username)

	if err != nil || !CheckPasswordHash(payload.Password, user.Password) {
		return ctx.String(http.StatusBadRequest, "invalid login")
	}

	tokens, err := GenerateJwtPair("s3cret", user.Username, 10, 600)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to generate token")
	}

	return ctx.JSON(http.StatusOK, tokens)
}

func (h *Handler) SignUp(ctx echo.Context) error {
	var payload RegistrationRequest
	err := ctx.Bind(&payload)
	if err != nil || payload.Username == "" || payload.Password == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	_, err = h.Provider.GetByUsername(payload.Username)

	if err == nil {
		return ctx.String(http.StatusBadRequest, "user already exists")
	}

	user, err := h.Provider.Create(user.User{
		Username: payload.Username,
		Password: MustHashPassword(payload.Password),
		Active:   false,
	})

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed creating user")
	}

	return ctx.JSON(http.StatusCreated, user)
}
