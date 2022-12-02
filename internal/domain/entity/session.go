package entity

import (
	"context"
)

type Session struct {
	Ctx context.Context //nolint:containedctx

	ID   string
	User User

	Transaction AbstractTransaction
}

func (sess Session) IsAuthorized() bool {
	return sess.User.ID != 0
}
