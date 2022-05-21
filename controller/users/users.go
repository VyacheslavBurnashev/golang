package users

import (
	"golang/domain/users"
	"golang/services"
	"golang/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewRequestError("invalid body")
		c.JSON(err.Status, err)
		return
	}
	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, result)
}
