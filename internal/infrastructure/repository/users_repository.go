package repository

import (
	"context"

	"github.com/ssummers02/booking/internal/domain/dto"
	"github.com/ssummers02/booking/internal/domain/entity"
	"github.com/ssummers02/booking/internal/infrastructure/dbmodel"

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

func (r *UserRepository) UpdateUser(ctx context.Context, e entity.User) (entity.User, error) {
	u := dto.UserToDB(e)

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Update("users").
			Set("first_name", u.FirstName).
			Set("surname", u.Surname).
			Set("middle_name", u.MiddleName).
			Set("phone", u.Phone).
			Where("id = ?", e.ID).
			Exec()

		if err != nil {
			return err
		}

		return err
	})

	return dto.UserFromDB(u), err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	return r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.DeleteFrom("users").
			Where("id = ?", id).
			Exec()

		return err
	})
}

func (r *UserRepository) GetUsersByIDs(ctx context.Context, ids []int64) ([]entity.User, error) {
	var u []dbmodel.User

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Select("*").
			From("users").
			Where("id IN ?", ids).
			Load(&u)

		return err
	})

	return dto.UsersFromDB(u), err
}
