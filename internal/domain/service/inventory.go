package service

import (
	"context"

	"github.com/ssummers02/booking/internal/domain"
	"github.com/ssummers02/booking/internal/domain/entity"
)

type InventoryService struct {
	repo InventoryStorage

	resortsService *ResortsService
}

func NewInventoryService(repo InventoryStorage, resortsService *ResortsService) *InventoryService {
	return &InventoryService{
		repo:           repo,
		resortsService: resortsService,
	}
}

func (s *InventoryService) GetInventoryByID(ctx context.Context, id int64) (entity.Inventory, error) {
	return s.repo.GetInventoryByID(ctx, id)
}

func (s *InventoryService) GetInventoryByResortID(ctx context.Context, resortID int64) ([]entity.Inventory, error) {
	return s.repo.GetInventoryByResortID(ctx, resortID)
}

func (s *InventoryService) CreateInventory(ctx context.Context, e entity.Inventory) (entity.Inventory, error) {
	user := ctx.Value("user").(entity.User)
	if !user.IsOwnerRole() {
		return entity.Inventory{}, domain.NewError(domain.ErrCodeForbidden, "user is not owner")
	}

	resort, err := s.resortsService.GetResortByID(ctx, e.ResortID)
	if err != nil {
		return entity.Inventory{}, err
	}

	if resort.OwnerID != user.ID {
		return entity.Inventory{}, domain.NewError(domain.ErrCodeForbidden, "user is not owner")
	}

	return s.repo.CreateInventory(ctx, e)
}

func (s *InventoryService) DeleteInventory(ctx context.Context, id int64) error {
	user, ok := ctx.Value("user").(entity.User)
	if !ok {
		return domain.NewError(domain.ErrCodeForbidden, "user is not role owner")
	}

	inventory, err := s.GetInventoryByID(ctx, id)
	if err != nil {
		return err
	}

	resort, err := s.resortsService.GetResortByID(ctx, inventory.ResortID)
	if err != nil {
		return err
	}

	if resort.OwnerID != user.ID {
		return domain.NewError(domain.ErrCodeForbidden, "user is not owner")
	}

	return s.repo.DeleteInventory(ctx, id)
}

func (s *InventoryService) UpdateInventory(ctx context.Context, e entity.Inventory) (entity.Inventory, error) {
	user, ok := ctx.Value("user").(entity.User)
	if !ok {
		return entity.Inventory{}, domain.NewError(domain.ErrCodeForbidden, "user is not role owner")
	}

	inventory, err := s.GetInventoryByID(ctx, e.ID)
	if err != nil {
		return entity.Inventory{}, err
	}

	resort, err := s.resortsService.GetResortByID(ctx, inventory.ResortID)
	if err != nil {
		return entity.Inventory{}, err
	}

	if resort.OwnerID != user.ID {
		return entity.Inventory{}, domain.NewError(domain.ErrCodeForbidden, "user is not owner")
	}

	return s.repo.UpdateInventory(ctx, e)
}

func (s *InventoryService) GetInventoriesTypes(ctx context.Context) ([]entity.InventoryType, error) {
	return s.repo.GetInventoriesTypes(ctx)
}

func (s *InventoryService) GetInventoriesByFilters(ctx context.Context, filters entity.InventoryFilter) ([]entity.Inventory, error) {
	return s.repo.GetInventoriesByFilters(ctx, filters)
}

func (s *InventoryService) UpdateImg(ctx context.Context, img entity.Img) (entity.Img, error) {
	user, ok := ctx.Value("user").(entity.User)
	if !ok {
		return entity.Img{}, domain.NewError(domain.ErrCodeForbidden, "user is not role owner")
	}

	inventory, err := s.GetInventoryByID(ctx, img.InventoryID)
	if err != nil {
		return entity.Img{}, err
	}

	resort, err := s.resortsService.GetResortByID(ctx, inventory.ResortID)
	if err != nil {
		return entity.Img{}, err
	}

	if resort.OwnerID != user.ID {
		return entity.Img{}, domain.NewError(domain.ErrCodeForbidden, "user is not owner")
	}

	return s.repo.UpdateImg(ctx, img)
}

func (s *InventoryService) GetImgByInventoryID(ctx context.Context, id int64) (entity.Img, error) {
	return s.repo.GetImgByInventoryID(ctx, id)
}
