package movie

import (
	"net/http"

	"github.com/abrarnaim015/KP-golang-tgs8/configs"
	basemodel "github.com/abrarnaim015/KP-golang-tgs8/models/base"
	moviemodel "github.com/abrarnaim015/KP-golang-tgs8/models/movie"
	"github.com/labstack/echo/v4"
)

func GetMoviesController(c echo.Context) error {
	var movie []moviemodel.Movie
	result := configs.DB.Find(&movie)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, basemodel.BaseResponse{
			Status: false,
			Message: "Failed get data from database",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, basemodel.BaseResponse{
		Status: true,
		Message: "Success",
		Data: movie,
	})
}