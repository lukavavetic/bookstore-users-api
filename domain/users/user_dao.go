package users

import (
	"fmt"
	"github.com/lukavavetic/bookstore-users-api/datasource/mysql/users_db"
	"github.com/lukavavetic/bookstore-users-api/utils/date"
	"github.com/lukavavetic/bookstore-users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	querySelectUser = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(querySelectUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	defer stmt.Close()

	res := stmt.QueryRow(user.Id)

	if err := res.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		fmt.Println(err)
		return errors.InternalServerError(err.Error())
	}


	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return errors.InternalServerError(saveErr.Error())
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	user.Id = userId

	return nil
}