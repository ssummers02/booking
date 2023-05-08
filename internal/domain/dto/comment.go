package dto

import (
	"github.com/ssummers02/booking/internal/delivery/api/restmodel"
	"github.com/ssummers02/booking/internal/domain/entity"
	"github.com/ssummers02/booking/internal/infrastructure/dbmodel"
)

func CommentFromDB(comment dbmodel.Comment) entity.Comment {
	return entity.Comment{
		ID:          comment.ID,
		UserID:      comment.UserID,
		InventoryID: comment.InventoryID,
		Text:        comment.Text,
		Rating:      comment.Rating,
		CreatedAt:   comment.CreatedAt,
	}
}

func CommentToDB(comment entity.Comment) dbmodel.Comment {
	return dbmodel.Comment{
		UserID:      comment.UserID,
		InventoryID: comment.InventoryID,
		Text:        comment.Text,
		Rating:      comment.Rating,
	}
}

func CommentsFromDB(comments []dbmodel.Comment) []entity.Comment {
	result := make([]entity.Comment, 0, len(comments))

	for i := range comments {
		result = append(result, CommentFromDB(comments[i]))
	}

	return result
}

func CommentToRest(comment entity.Comment) restmodel.Comment {
	return restmodel.Comment{
		ID:          comment.ID,
		UserID:      comment.UserID,
		InventoryID: comment.InventoryID,
		Rating:      comment.Rating,
		Text:        comment.Text,
		CreatedAt:   comment.CreatedAt,
	}
}
func CommentsToRest(comments []entity.Comment) []restmodel.Comment {
	result := make([]restmodel.Comment, 0, len(comments))

	for i := range comments {
		result = append(result, CommentToRest(comments[i]))
	}

	return result
}

func CommentCreateEntity(comment restmodel.Comment, userID int64) entity.Comment {
	return entity.Comment{
		UserID:      userID,
		InventoryID: comment.InventoryID,
		Text:        comment.Text,
		Rating:      comment.Rating,
	}
}
