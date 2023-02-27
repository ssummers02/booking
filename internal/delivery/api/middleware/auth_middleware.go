package middleware

import (
	"booking/internal/domain"
	"booking/internal/domain/service"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	users service.UserStorage
}

type errResponseData struct {
	IsOK    bool        `json:"is_ok"`
	Payload interface{} `json:"payload"`
}

type ErrResponse struct {
	Data errResponseData `json:"data"`
}

func NewAuth(users service.UserStorage) *Auth {
	return &Auth{
		users: users,
	}
}

// Handler creates a new callback that is run when we require authentication.
func (m *Auth) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")

			if tokenString != "" {
				mail, err := m.ValidateToken(tokenString[7:])
				if err != nil && mail == "" {
					data, _ := json.Marshal(ErrResponse{
						Data: errResponseData{
							IsOK:    false,
							Payload: domain.NewError(domain.ErrCodeNotAuthorized, "request does not contain an access token"),
						},
					})
					_, _ = w.Write(data)

					return
				}

				user, err := m.users.GetUsersByEmail(r.Context(), mail)
				if err != nil {
					if !errors.Is(err, domain.ErrNotFound) {
						next.ServeHTTP(w, r)

						return
					}
				}

				r = r.WithContext(context.WithValue(r.Context(), "user", user))
				log.Println("user", user)
			}

			next.ServeHTTP(w, r)
		},
	)
}

var jwtKey = []byte("supersecretkey") // env of file
type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (m *Auth) ValidateToken(signedToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")

		return "", err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")

		return "", err
	}

	return claims.Email, nil
}

func (m *Auth) GenerateJWT(email string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)

	return
}
