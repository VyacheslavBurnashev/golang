package users

import (
	"golang/utils/errors"
	"strings"
)

type User struct {
	ID       int64  `json: "ID`
	Username string `json: "username"`
	Password string `json: "password"`
	Email    string `json: "email"`
}

func (user *User) Validate() *errors.Errors {
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
	if user.Email == "" {
		return errors.NewRequestError("invalid address")
	}

	if user.Password == "" {
		return errors.NewRequestError("invalid password")
	}
	return nil
}
