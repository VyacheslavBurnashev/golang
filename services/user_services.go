package services

import (
	"golang/domain/users"
	"golang/utils/errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *errors.Errors) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	Slice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, errors.NewRequestError(("failed to encrypt the password"))
	}
	user.Password = string(Slice[:])

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
