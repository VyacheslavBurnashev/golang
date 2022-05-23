package users

import (
	"fmt"
	"golang/database/mysql/users_data"
	"golang/utils/errors"
)

var (
	InsertUser     = "INSERT INTO users (username, password, email) VALUES (?, ?, ?);"
	GetUserbyEmail = "SELECT id, username, password, email FROM users WHERE email=?;"
	GetUserbyid    = "SELECT id, username, password, email FROM users WHERE id=?;"
	ChangeEmail    = "UPDATE id, username, password, email FROM users WHERE email=?;"
)

func (user *User) Save() *errors.Errors {
	stmt, err := users_data.Client.Prepare(InsertUser)
	if err != nil {
		return errors.ServerError("database error1")
	}
	defer stmt.Close()

	insert, err := stmt.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		return errors.ServerError("database error2")
	}

	userID, err := insert.LastInsertId()
	if err != nil {
		return errors.ServerError("database error3")
	}
	user.ID = userID
	return nil
}

func (user *User) GetByEmail() *errors.Errors {
	stmt, err := users_data.Client.Prepare(GetUserbyEmail)
	if err != nil {
		return errors.ServerError("invalid email")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email)

	if geterr := result.Scan(&user.ID, &user.Username, &user.Email, &user.Password); geterr != nil {
		fmt.Println(geterr)
		return errors.ServerError("database error5")
	}
	return nil
}

func (user *User) GetUserByID() *errors.Errors {
	stmt, err := users_data.Client.Prepare(GetUserbyid)
	if err != nil {
		return errors.NewRequestError("database error6")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.Username, &user.Email); getErr != nil {
		return errors.NewRequestError("database error7")
	}
	return nil
}
