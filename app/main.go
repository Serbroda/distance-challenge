package main

import (
	"fmt"
	"github.com/Serbroda/distance-challenge/db"
	_ "github.com/Serbroda/distance-challenge/docs"
	"github.com/Serbroda/distance-challenge/handlers"
	"github.com/Serbroda/distance-challenge/security"
	"github.com/Serbroda/distance-challenge/services"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Distance Challenge API
// @version 1.0
// @description Distance Challenge API.

// @host localhost
// @BasePath /
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @scheme bearer
// @bearerFormat: JWT
func main() {
	db := db.Connect(":memory:")
	us := services.UserService{DB: db}
	//rs := run.RunService{DB: db}

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/api/docs/*", echoSwagger.WrapHandler)

	handlers.RegisterAuthHandlers(e, &handlers.AuthHandler{UserService: &us}, "")

	api := e.Group("/api/v1")
	api.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(security.Claims)
		},
		SigningKey: []byte("s3cret"),
	}))
	handlers.RegisterUserHandlers(api, &handlers.UserHandler{UserService: &us}, "")

	for _, r := range e.Routes() {
		fmt.Printf("%s %s\n", r.Method, r.Path)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
