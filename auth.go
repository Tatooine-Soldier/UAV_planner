package main

import (
	"context"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register(ctx context.Context) error {
	var data User

	_, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if err != nil {
		return http.ErrNotSupported
	}

	user := &User{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}

	err = CreateUser(user)
	return err
}
