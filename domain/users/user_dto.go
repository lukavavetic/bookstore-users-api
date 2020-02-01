package users

import (
	"github.com/lukavavetic/bookstore-users-api/utils/errors"
	"strings"
)

type User struct {
	Id 			int64   `json:"id"`
	FirstName 	string  `json:"first_name"`
	LastName 	string  `json:"last_name"`
	Email     	string  `json:"email"`
	DateCreated string  `json:"date_created"`
}

// this is a function
//func Validate(user *User) *errors.RestErr {
//	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
//
//	if user.Email == "" {
//		return errors.BadRequest("Invalid email address!")
//	}
//
//	return nil
//}

// this is a method
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.BadRequest("Invalid email address!")
	}

	return nil
}