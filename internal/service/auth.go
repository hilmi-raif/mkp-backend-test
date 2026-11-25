package service

import (
	"context"
	"errors"
	"time"

	"mkp-backend-test/internal/dto"
	"mkp-backend-test/internal/helper"
	"mkp-backend-test/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Auth struct {
	users  *repository.Postgres
	secret []byte
}

func NewAuth(users *repository.Postgres, secret []byte) *Auth {
	return &Auth{users: users, secret: secret}
}

func (s *Auth) Login(ctx context.Context, input dto.Login) (string, error) {
	user, err := s.users.FindUser(ctx, input.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", helper.ErrCredentials
	}
	if err != nil {
		return "", err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)) != nil {
		return "", helper.ErrCredentials
	}
	return helper.SignToken(s.secret, user.ID, time.Now().Add(time.Hour))
}
