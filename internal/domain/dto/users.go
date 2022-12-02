package dto

import (
	"booking/internal/delivery/api/restmodel"
	"booking/internal/domain/entity"
	"booking/internal/infrastructure/dbmodel"
)

func UserFromDB(u dbmodel.User) entity.User {
	return entity.User{
		ID:         u.ID,
		FirstName:  u.FirstName,
		Surname:    u.Surname,
		MiddleName: u.MiddleName,
		Email:      u.Email,
		Password:   u.Password,
		Phone:      u.Phone,
		RoleID:     entity.RoleType(u.RoleID),
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

func UserToDB(e entity.User) dbmodel.User {
	return dbmodel.User{
		ID:         e.ID,
		FirstName:  e.FirstName,
		Surname:    e.Surname,
		MiddleName: e.MiddleName,
		Email:      e.Email,
		Password:   e.Password,
		Phone:      e.Phone,
		RoleID:     int64(e.RoleID),
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
	}
}

func UserFromRest(r restmodel.User) entity.User {
	return entity.User{
		ID:         r.ID,
		FirstName:  r.FirstName,
		Surname:    r.Surname,
		MiddleName: r.MiddleName,
		Email:      r.Email,
		Password:   r.Password,
		Phone:      r.Phone,
		RoleID:     entity.RoleType(r.RoleID),
	}
}
func UserToRest(e entity.User, token string) restmodel.User {
	return restmodel.User{
		ID:         e.ID,
		FirstName:  e.FirstName,
		Surname:    e.Surname,
		MiddleName: e.MiddleName,
		Email:      e.Email,
		Phone:      e.Phone,
		RoleID:     int64(e.RoleID),
		Token:      token,
	}
}
