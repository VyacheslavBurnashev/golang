package users

import (
	"golang/domain/users"
	"golang/services"
	"golang/utils/errors"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func Login(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewRequestError("invalid body")
		c.JSON(err.Status, err)
		return
	}
	result, loginerr := services.GetUser(user)
	if loginerr != nil {
		c.JSON(loginerr.Status, loginerr)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(result.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	})

	token, err := claims.SignedString([]byte(""))
	if err != nil {
		err := errors.NewRequestError("login failed")
		c.JSON(err.Status, err)
		return
	}
	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, result)
}

func Get(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		getErr := errors.NewRequestError("can't import cookie")
		c.JSON(getErr.Status, getErr)
		return
	}
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	if err != nil {
		restErr := errors.NewRequestError("can't parse the cookie")
		c.JSON(restErr.Status, restErr)
		return
	}
	claims := token.Claims.(*jwt.StandardClaims)
	issuer, err := strconv.ParseInt(claims.Issuer, 10, 64)
	if err != nil {
		restErr := errors.NewRequestError("id should be int")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, restErr := services.GetUserByID(issuer)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, result)

}
func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
