package main

import (
	"github.com/Serbroda/distance-challenge/db"
	_ "github.com/Serbroda/distance-challenge/docs"
	"github.com/Serbroda/distance-challenge/security"
	"github.com/Serbroda/distance-challenge/user"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /
// @schemes http
func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/api/docs/*", echoSwagger.WrapHandler)

	db := db.Connect(":memory:")
	us := user.UserServiceGorm{DB: db}

	security.RegisterHandlers(e, &security.Handler{Provider: &us}, "")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(security.Claims)
		},
		SigningKey: []byte("s3cret"),
	}

	api := e.Group("/api/v1")
	api.Use(echojwt.WithConfig(config))

	e.Logger.Fatal(e.Start(":1323"))
}
