package auth

import (
	"net/http"

	"github.com/abrarnaim015/KP-golang-tgs8/configs"
	"github.com/abrarnaim015/KP-golang-tgs8/middleware"
	"github.com/abrarnaim015/KP-golang-tgs8/models/base"
	usermodel "github.com/abrarnaim015/KP-golang-tgs8/models/user"
	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error{
	var user usermodel.User
	c.Bind(&user)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hashPassword)

	// result := configs.DB.Find(&user)
	result := configs.DB.Where("email = ?", user.Email).First(&user)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, base.BaseResponse{
			Status: false,
			Message: "User Not Found",
			Data: nil,
		})
	}

	var authResponse = usermodel.ResponseAuth {
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
		Token: middleware.GenerateTokenJWT(user.Id, user.Name),
	}
	
	
	return c.JSON(http.StatusOK, base.BaseResponse{
		Status: true,
		Message: "Success Loggin",
		Data: authResponse,
	})
}