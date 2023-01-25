package main

import (
	"github.com/Serbroda/distance-challenge/db"
	"github.com/Serbroda/distance-challenge/security"
	"github.com/Serbroda/distance-challenge/user"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	db := db.Connect(":memory:")
	us := user.UserServiceGorm{DB: db}

	security.RegisterHandlers(e, &security.Handler{Provider: &us})

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(security.Claims)
		},
		SigningKey: []byte("s3cret"),
	}

	e.Use(echojwt.WithConfig(config))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
