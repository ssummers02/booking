package service

import (
	"context"

	"github.com/ssummers02/booking/internal/domain"
	"github.com/ssummers02/booking/internal/domain/entity"
)

type CommentService struct {
	repo CommentStorage
}

func NewCommentService(repo CommentStorage) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (s *CommentService) GetCommentByID(ctx context.Context, id int64) (entity.Comment, error) {
	Comment, err := s.repo.GetCommentByID(ctx, id)
	if err != nil {
		return entity.Comment{}, err
	}

	return Comment, nil
}

func (s *CommentService) CreateComment(ctx context.Context, comments entity.Comment) (entity.Comment, error) {
	createComment, err := s.repo.CreateComment(ctx, comments)
	if err != nil {
		return entity.Comment{}, err
	}

	return s.GetCommentByID(ctx, createComment.ID)
}

func (s *CommentService) GetCommentByResortID(ctx context.Context, resortID int64) ([]entity.Comment, error) {
	comments, err := s.repo.GetCommentsByResort(ctx, resortID)
	if err != nil {
		return []entity.Comment{}, err
	}

	return comments, nil
}

func (s *CommentService) GetCommentsByInventory(ctx context.Context, inventoryID int64) ([]entity.Comment, error) {
	comments, err := s.repo.GetCommentsByInventory(ctx, inventoryID)
	if err != nil {
		return []entity.Comment{}, err
	}

	return comments, nil
}

func (s *CommentService) DeleteComment(ctx context.Context, id int64) error {
	user, ok := ctx.Value("user").(entity.User)
	if !ok {
		return domain.NewError(domain.ErrCodeAlreadyExists, "user is not authorized")
	}

	commentByID, err := s.GetCommentByID(ctx, id)
	if err != nil {
		return err
	}

	if commentByID.UserID != user.ID {
		return domain.NewError(domain.ErrCodeForbidden, "user is not owner of comment")
	}

	err = s.repo.DeleteCommentsByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
