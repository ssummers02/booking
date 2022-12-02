package dto

import (
	"booking/internal/domain/entity"
	"context"
)

type ctxKey string

const CtxSession ctxKey = "session"

func Session(ctx context.Context) entity.Session {
	value := ctx.Value(CtxSession)

	sess, ok := value.(entity.Session)
	if !ok {
		return entity.Session{}
	}

	return sess
}
