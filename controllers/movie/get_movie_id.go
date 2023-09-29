package movie

import (
	"net/http"

	"github.com/abrarnaim015/KP-golang-tgs8/configs"
	"github.com/abrarnaim015/KP-golang-tgs8/models/base"
	moviemodel "github.com/abrarnaim015/KP-golang-tgs8/models/movie"
	"github.com/labstack/echo/v4"
)

func GetMovieById(c echo.Context) error {
	var movie moviemodel.Movie
	id := c.Param("id")

	result := configs.DB.Where("id = ?", id).First(&movie)
	if result.Error != nil {
			return c.JSON(http.StatusNotFound, base.BaseResponse{
					Status:  false,
					Message: "Movie not exists",
					Data:    nil,
			})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
			Status:  true,
			Message: "Success get Movie",
			Data:    movie,
	})
}
