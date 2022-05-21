package users

import (
	"golang/database/mysql/users_data"
	"golang/utils/errors"
)

var (
	InsertUser = "INSERT INTO users (username, password, email) VALUES (?, ?, ?);"
)

func (user *User) Save() *errors.Errors {
	stmt, err := users_data.Client.Prepare(InsertUser)
	if err != nil {
		return errors.NewRequestError("database error")
	}
	defer stmt.Close()

	insert, err := stmt.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		return errors.NewRequestError("database error")
	}

	userID, err := insert.LastInsertId()
	if err != nil {
		return errors.NewRequestError("database error")
	}
	user.ID = userID
	return nil
}
