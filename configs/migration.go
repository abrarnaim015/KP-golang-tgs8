package configs

import (
	"github.com/abrarnaim015/KP-golang-tgs8/models/movie"
	"github.com/abrarnaim015/KP-golang-tgs8/models/user"
)

func initMigrate(){
	DB.AutoMigrate(&user.User{})
	DB.AutoMigrate(&movie.Movie{})
}