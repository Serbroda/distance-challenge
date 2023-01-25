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
	Username string
	Password string
}

type TokenPair struct {
	AccessToken  Jwt
	RefreshToken Jwt
}

func RegisterHandlers(e *echo.Echo, h *Handler) {
	e.POST("/login", h.Login)
}

func (h *Handler) Login(ctx echo.Context) error {
	var payload LoginRequest
	err := ctx.Bind(&payload)
	if err != nil || payload.Username == "" || payload.Password == "" {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	user, err := h.Provider.GetByUsername(payload.Username)

	if !CheckPasswordHash(user.Password, payload.Password) {
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
