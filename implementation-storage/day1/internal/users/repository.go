package users

import (
	"errors"
)

var (
	ErrNotFound = errors.New("user not found in the given repository")
	ErrAlreadyExists = errors.New("user already exists in the giver repository")
)

type Repository interface{
	GetById(id int) (*User, error)
	Create(user *User ) error
	Update (user *User) error
	Delete(id int) error
}