package user

import (
	"net/http"

	"github.com/abrarnaim015/KP-golang-tgs8/configs"
	basemodel "github.com/abrarnaim015/KP-golang-tgs8/models/base"
	usermodel "github.com/abrarnaim015/KP-golang-tgs8/models/user"

	"github.com/labstack/echo/v4"
)


func GetUsersController(c echo.Context) error{
	var users []usermodel.User
	result := configs.DB.Find(&users)
	
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, basemodel.BaseResponse{
			Status: false,
			Message: "Failed get data from database",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, basemodel.BaseResponse{
		Status: true,
		Message: "Success",
		Data: users,
	})
}
