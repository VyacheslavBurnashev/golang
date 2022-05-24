package users

import (
	"golang/domain/users"
	"golang/forms"
	"golang/services"
	"golang/utils/errors"
	"net/http"
	"strconv"

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
	var data forms.LoginUserCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}
	result, err := services.GetUserByEmail(data.Email)
	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User is not found"})
		c.Abort()
		return
	}
	if !result.IsVerified {
		c.JSON(403, gin.H{"message": "Account is not verified"})
		c.Abort()
		return
	}
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem logging into your account"})
		c.Abort()
		return
	}

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

func PasswordReset(c *gin.Context) {
	var data forms.PasswordResetCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()
		return
	}
	if data.Password != data.Confirm {
		c.JSON(400, gin.H{"message": "Password don't match"})
	}
	resetToken, _ := c.GetQuery("token")
	userID, _ := services.DecodeNonAuthToken(resetToken)
	result, err := services.GetUserByEmail(userID)

	if err != nil {
		c.JSON(500, gin.H{"message": "Something wrong happened"})
		c.Abort()
		return
	}
	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User was not found"})
		c.Abort()
		return
	}

}

func ResetLink(c *gin.Context) {
	var data forms.ResendCommand
	if (c.BindJSON(&data)) != nil {
		c.JSON(400, gin.H{"message": "Provided all fields"})
		c.Abort()
		return
	}
	result, err := services.GetUserByEmail(data.Email)
	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User was not found"})
		c.Abort()
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"message": "Wrong happened"})
		c.Abort()
		return
	}

	resetToken, _ := services.DecodeNonAuthToken(result.Email)
	link := resetToken
	body := link
	html := body
	email := services.SendMail("Reset Password", body, result.Email, html, result.Username)

	if email == true {
		c.JSON(200, gin.H{"message": "Check email"})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"message": "Issue"})
		c.Abort()
	}
}

func VerifyAccount(c *gin.Context) {
	verifyToken, _ := c.GetQuery("")
	userID, _ := services.DecodeNonAuthToken(verifyToken)
	result, err := services.GetUserByEmail(userID)
	if err != nil {
		c.JSON(500, gin.H{"message": "Something wrong happened"})
		c.Abort()
		return
	}
	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User was not found"})
		c.Abort()
		return
	}

}
