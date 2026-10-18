package middleware

import (
	"fmt"
	"golink/internal/config"
	"net/http"
	"uuid"

	"github.com/golang-jwt/jwt/v5"
)

type MiddlewareInterface interface {
	AuthUser(next http.HandlerFunc) http.HandlerFunc
}

type Middleware struct {
	env *config.ENV
}

func NewMiddleware(env *config.ENV) *Middleware {
	return &Middleware{env: env}
}

func validarToken(tokenString, secretKey string) (*uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("token inválido ou expirado: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("erro ao mapear os claims do token")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, fmt.Errorf("campo 'sub' não encontrado ou inválido no token")
	}

	id, err := uuid.Parse(sub)
	if err != nil {
		return nil, fmt.Errorf("Erro ao retornar id")
	}

	return &id, nil
}
