package routes

import (
	"os"

	authcontroller "github.com/abrarnaim015/KP-golang-tgs8/controllers/auth"
	moviecontroller "github.com/abrarnaim015/KP-golang-tgs8/controllers/movie"
	usercontroller "github.com/abrarnaim015/KP-golang-tgs8/controllers/user"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.POST("/login", authcontroller.LoginController)
	e.POST("/register", authcontroller.RegisterController)

	eAuthUser := e.Group("")
	eAuthUser.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))
	eAuthUser.GET("/users", usercontroller.GetUsersController)

	eAuthUser.POST("/movie", moviecontroller.CreateMovie)
	eAuthUser.GET("/movie", moviecontroller.GetMoviesController)
	eAuthUser.GET("/movie/:id", moviecontroller.GetMovieById)

}