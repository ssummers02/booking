package service

import (
	"booking/internal/domain/entity"
	"context"
)

type UsersService struct {
	repo UserStorage
}

func NewUsersService(repo UserStorage) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) CreateUser(ctx context.Context, e entity.User) (entity.User, error) {
	if err := e.HashPassword(e.Password); err != nil {
		return entity.User{}, err
	}

	return s.repo.CreateUser(ctx, e)
}

func (s *UsersService) GetUsersByEmail(ctx context.Context, mail string) (entity.User, error) {
	return s.repo.GetUsersByEmail(ctx, mail)
}
