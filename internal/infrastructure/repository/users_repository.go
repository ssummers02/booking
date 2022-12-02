package repository

import (
	"booking/internal/domain/dto"
	"booking/internal/domain/entity"
	"booking/internal/infrastructure/dbmodel"
	"context"

	"github.com/gocraft/dbr/v2"
)

type UserRepository struct {
	*DBConn
}

func NewUserRepository(db *DBConn) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetUsersByEmail(ctx context.Context, mail string) (entity.User, error) {
	var u dbmodel.User

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		return tx.Select("*").
			From("users").
			Where("email = ?", mail).
			LoadOne(&u)
	})

	return dto.UserFromDB(u), err
}

func (r *UserRepository) CreateUser(ctx context.Context, e entity.User) (entity.User, error) {
	u := dto.UserToDB(e)
	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		err := tx.InsertInto("users").
			Returning("id").
			Columns("first_name", "surname", "middle_name", "email", "password", "phone", "role_id").
			Record(&u).
			Load(&u.ID)

		if err != nil {
			return err
		}

		return err
	})

	if err != nil {
		return entity.User{}, err
	}

	return dto.UserFromDB(u), err
}
