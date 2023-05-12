package service

import (
	"context"

	"github.com/ssummers02/booking/internal/domain/entity"
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

func (s *UsersService) UpdateUser(ctx context.Context, e entity.User) (entity.User, error) {
	return s.repo.UpdateUser(ctx, e)
}

func (s *UsersService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.DeleteUser(ctx, id)
}
func (s *UsersService) GetUsersByIDs(ctx context.Context, ids []int64) (map[int64]entity.User, error) {
	users, err := s.repo.GetUsersByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	userMap := make(map[int64]entity.User, len(users))
	for _, user := range users {
		userMap[user.ID] = user
	}
	return userMap, nil
}
func (s *UsersService) GetUsersByID(ctx context.Context, id int64) (entity.User, error) {
	ds, err := s.GetUsersByIDs(ctx, []int64{id})
	if err != nil {
		return entity.User{}, err
	}
	return ds[id], nil
}
