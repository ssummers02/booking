package service

import (
	"booking/internal/domain"
	"booking/internal/domain/entity"
	"context"
)

type ResortsService struct {
	repo ResortStorage
}

func NewResortsService(repo ResortStorage) *ResortsService {
	return &ResortsService{repo: repo}
}

func (s *ResortsService) GetResorts(ctx context.Context) ([]entity.Resort, error) {
	return s.repo.GetResorts(ctx)
}

func (s *ResortsService) GetResortByID(ctx context.Context, id int64) (entity.Resort, error) {
	return s.repo.GetResortByID(ctx, id)
}

func (s *ResortsService) CreateResort(ctx context.Context, e entity.Resort) (entity.Resort, error) {
	user := ctx.Value("user").(entity.User)
	if !user.IsOwnerRole() {
		return entity.Resort{}, domain.NewError(domain.ErrCodeForbidden, "user is not owner")
	}

	e.OwnerID = user.ID

	return s.repo.CreateResort(ctx, e)
}

func (s *ResortsService) DeleteResort(ctx context.Context, id int64) error {
	user, ok := ctx.Value("user").(entity.User)
	if !ok {
		return domain.NewError(domain.ErrCodeForbidden, "user is not role owner")
	}

	resort, err := s.GetResortByID(ctx, id)
	if err != nil {
		return err
	}

	if resort.OwnerID != user.ID {
		return domain.NewError(domain.ErrCodeForbidden, "user is not owner")
	}

	return s.repo.DeleteResort(ctx, id)
}

func (s *ResortsService) UpdateResort(ctx context.Context, e entity.Resort) (entity.Resort, error) {
	user, ok := ctx.Value("user").(entity.User)
	if !ok {
		return entity.Resort{}, domain.NewError(domain.ErrCodeForbidden, "user is not owner")
	}

	resort, err := s.GetResortByID(ctx, e.ID)
	if err != nil {
		return entity.Resort{}, err
	}

	if resort.OwnerID != user.ID {
		return entity.Resort{}, domain.NewError(domain.ErrCodeForbidden, "user is not owner")
	}

	return s.repo.UpdateResort(ctx, e)
}

func (s *ResortsService) GetResortsByFilter(ctx context.Context, filter entity.Filter) ([]entity.Resort, error) {
	return s.repo.GetResortsByFilter(ctx, filter)
}

func (s *ResortsService) GetCities(ctx context.Context) ([]entity.City, error) {
	return s.repo.GetCities(ctx)
}

func (s *ResortsService) CheckReservation(ctx context.Context, e entity.Booking) error {
	return s.repo.CheckReservation(ctx, e)
}
