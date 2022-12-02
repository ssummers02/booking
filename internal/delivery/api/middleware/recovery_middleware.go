package middleware

import (
	"log"
	"net/http"
)

type Recovery struct{}

func NewRecovery() *Recovery {
	return &Recovery{}
}

// Handler catches the panic.
func (r *Recovery) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			msg := recover()
			if msg != nil {
				log.Println("recover: ", msg)

				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
