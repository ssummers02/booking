package repository

import (
	"context"
	"errors"

	"github.com/ssummers02/booking/internal/domain"
	"github.com/ssummers02/booking/internal/domain/dto"
	"github.com/ssummers02/booking/internal/domain/entity"
	"github.com/ssummers02/booking/internal/infrastructure/dbmodel"

	"github.com/gocraft/dbr/v2"
)

type CommentsRepository struct {
	*DBConn
}

func NewCommentsRepository(db *DBConn) *CommentsRepository {
	return &CommentsRepository{db}
}
func (r *CommentsRepository) CreateComment(ctx context.Context, booking entity.Comment) (entity.Comment, error) {
	commentToDB := dto.CommentToDB(booking)

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		return tx.InsertInto("comments").
			Returning("id").
			Columns("user_id", "inventory_id", "rating", "text").
			Record(&commentToDB).
			Load(&commentToDB.ID)
	})

	return dto.CommentFromDB(commentToDB), err
}
func (r *CommentsRepository) GetCommentByID(ctx context.Context, id int64) (entity.Comment, error) {
	var comment dbmodel.Comment

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		return tx.Select("comments.*, users.surname").
			From("comments").
			LeftJoin("users", "comments.user_id = users.id").
			Where("comments.id = ?", id).
			LoadOne(&comment)
	})

	return dto.CommentFromDB(comment), err
}

func (r *CommentsRepository) GetCommentsByResort(ctx context.Context, id int64) ([]entity.Comment, error) {
	var comments []dbmodel.Comment

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Select("comments.*, users.surname").
			From("comments").
			LeftJoin("users", "comments.user_id = users.id").
			LeftJoin("inventory", "comments.inventory_id = inventory.id").
			Where("inventory.resort_id = ?", id).
			Load(&comments)

		return err
	})

	if errors.Is(err, domain.ErrNotFound) {
		return []entity.Comment{}, nil
	}

	return dto.CommentsFromDB(comments), err
}

func (r *CommentsRepository) GetCommentsByInventory(ctx context.Context, id int64) ([]entity.Comment, error) {
	var comments []dbmodel.Comment

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Select("comments.*").
			From("comments").
			Where("comments.inventory_id = ?", id).
			Load(&comments)

		return err
	})

	if errors.Is(err, domain.ErrNotFound) {
		return []entity.Comment{}, nil
	}

	return dto.CommentsFromDB(comments), err
}

func (r *CommentsRepository) DeleteCommentsByID(ctx context.Context, id int64) error {
	return r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.DeleteFrom("comments").
			Where("id = ?", id).
			Exec()

		return err
	})
}
