package entity

import (
	"errors"
	"time"
)

type User struct {
	Id        string
	Name      string
	Username  string
	Password  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(
	id string,
	name string,
	username string,
	password string,
	isActive bool,
) *User {
	return &User{
		Id:        id,
		Name:      name,
		Username:  username,
		Password:  password,
		IsActive:  isActive,
		CreatedAt: time.Now(),
	}
}

func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("name must not be empty")
	}

	if u.Username == "" {
		return errors.New("username must not be empty")
	}

	return nil
}
