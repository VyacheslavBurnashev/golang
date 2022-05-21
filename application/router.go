package application

import (
	"golang/controller/users"
	"os/user"
)

func mapUrls() {
	router.POST("api/register", users.Register)
	router.POST("api/login", user.Login)
	router.GET("api/user", user.Get)
	router.GET("api/logout", user.Logout)
}
