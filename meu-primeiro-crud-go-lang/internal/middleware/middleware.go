package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"meu-primeiro-crud-go-lang/internal/config"
	errapp "meu-primeiro-crud-go-lang/internal/config/err-app"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type Middleware struct {
	cfg config.Config
}

func NewMiddleware(cfg *config.Config) *Middleware {
	return &Middleware{cfg: *cfg}
}

func (m *Middleware) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		cookie, err := r.Cookie("session_token_user")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				errApp := errapp.NewUnauthorizedError("Token de sessão não fornecido", err)
				w.WriteHeader(errApp.Code)
				json.NewEncoder(w).Encode(errApp)
				return
			}
			errApp := errapp.NewInternalServerError("Erro do servidor", err)
			w.WriteHeader(errApp.Code)
			json.NewEncoder(w).Encode(errApp)
			return
		}

		tokenString := cookie.Value
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(m.cfg.Env.JWT_SECRET_KEY), nil
		})
		if err != nil || !token.Valid {
			errApp := errapp.NewUnauthorizedError("Token inválido ou expirado", err)
			w.WriteHeader(errApp.Code)
			json.NewEncoder(w).Encode(errApp)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			errApp := errapp.NewUnauthorizedError("Problema ao ler o token", err)
			w.WriteHeader(errApp.Code)
			json.NewEncoder(w).Encode(errApp)
			return
		}

		sub, ok := claims["sub"].(string)
		if !ok {
			errApp := errapp.NewUnauthorizedError("Identificador não encontrado", err)
			w.WriteHeader(errApp.Code)
			json.NewEncoder(w).Encode(errApp)
			return
		}

		ctx := context.WithValue(r.Context(), "user_sub", sub)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
