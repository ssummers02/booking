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
