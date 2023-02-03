package middleware

import (
	"log"
	"net/http"
)

type Cors struct{}

func NewCors() *Cors {
	return &Cors{}
}

// Handler creates a new callback that is run before handling any requests.
func (m *Cors) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			header := w.Header()

			header.Set("Accept", "application/json")
			header.Set("Content-Type", "application/json")
			header.Set("Access-Control-Allow-Origin", "*")
			header.Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")

			log.Printf("CORS: %s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		},
	)
}
