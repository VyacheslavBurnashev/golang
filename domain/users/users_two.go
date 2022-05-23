package users

import (
	"golang/utils/errors"
	"strings"
)

type User struct {
	ID         int64  `json:"_id,omitempty" bson:"_id,omitempty"`
	Username   string `json:"name" bson:"name"`
	Email      string `json:"email" bson:"email"`
	Password   string `json:"password" bson:"password"`
	IsVerified bool   `json:"is_verified" bson:"is_verified"`
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
