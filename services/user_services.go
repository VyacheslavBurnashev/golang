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

func GetUser(user users.User) (*users.User, *errors.Errors) {
	result := &users.User{Email: user.Email}

	if err := result.GetByEmail(); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return nil, errors.NewRequestError("failed to decrypt the password")
	}

	resultW := &users.User{ID: result.ID, Username: result.Username, Email: result.Email}
	return resultW, nil
}

func GetUserByID(userId int64) (*users.User, *errors.Errors) {
	result := &users.User{ID: userId}
	if err := result.GetUserByID(); err != nil {
		return nil, err

	}
	return result, nil
}
