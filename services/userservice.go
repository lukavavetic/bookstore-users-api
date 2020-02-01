package services

import (
	"github.com/lukavavetic/bookstore-users-api/domain"
	"github.com/lukavavetic/bookstore-users-api/utils/errors"
)

func CreateUser(user domain.User) (*domain.User, *errors.RestErr) {
	return &user, nil
}