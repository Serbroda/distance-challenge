package security

import (
	"github.com/Serbroda/distance-challenge/user"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Jwt = string

type Handler struct {
	Provider user.UserService
}

type Claims struct {
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RegistrationRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type TokenPair struct {
	AccessToken  Jwt `json:"access_token"`
	RefreshToken Jwt `json:"refresh_token"`
}

func RegisterHandlers(e *echo.Echo, h *Handler) {
	e.POST("/login", h.Login)
	e.POST("/register", h.Register)
}

func (h *Handler) Login(ctx echo.Context) error {
	var payload LoginRequest
	err := ctx.Bind(&payload)
	if err != nil || payload.Username == "" || payload.Password == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	user, err := h.Provider.GetByUsername(payload.Username)

	if err != nil || !CheckPasswordHash(payload.Password, user.Password) {
		return ctx.String(http.StatusBadRequest, "invalid login")
	}

	t, rt, err := generateTokenPair(user, []byte("s3cret"), 10, 600)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to generate token")
	}

	return ctx.JSON(http.StatusOK, TokenPair{
		AccessToken:  t,
		RefreshToken: rt,
	})
}

func (h *Handler) Register(ctx echo.Context) error {
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

func generateTokenPair(user user.User, secret []byte, accessTokenMinues, refreshTokenMinutes int64) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(accessTokenMinues))),
		},
	})
	t, err := token.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(refreshTokenMinutes))),
		},
	})
	rt, err := refreshToken.SignedString(secret)
	if err != nil {
		return "", "", err
	}
	return t, rt, nil
}
