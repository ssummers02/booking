package middleware

import (
	"booking/internal/domain/service"
)

// M provides list of all middlewares.
type M struct {
	Auth     *Auth
	Cors     *Cors
	Recovery *Recovery
}

func NewMiddlewares(i service.UserStorage) *M {
	return &M{
		Auth:     NewAuth(i),
		Cors:     NewCors(),
		Recovery: NewRecovery(),
	}
}
