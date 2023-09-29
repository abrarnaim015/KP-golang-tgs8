package movie

import (
	"net/http"

	"github.com/abrarnaim015/KP-golang-tgs8/configs"
	"github.com/abrarnaim015/KP-golang-tgs8/models/base"
	moviemodel "github.com/abrarnaim015/KP-golang-tgs8/models/movie"
	"github.com/labstack/echo/v4"
) 

func CreateMovie(c echo.Context) error {
	var movie moviemodel.Movie
	c.Bind(&movie)
	
	if movie.Title == "" {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status: false,
			Message: "Invalid Input Data",
			Data: nil,
		})
	}

	movieCheck := configs.DB.Where("title = ?", movie.Title).First(&movie)
	if movieCheck.Error == nil {
		return c.JSON(http.StatusNotFound, base.BaseResponse{
			Status: false,
			Message: "Movie already exists",
			Data: nil,
		})
	}

	result := configs.DB.Create(&movie)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status: false,
			Message: "Failed add data to database",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status: true,
		Message: "Success register Movie",
		Data: movie,
	})
}