package middleware

import (
	"backend/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type (
	Middleware interface {
		HandleUser() gin.HandlerFunc
	}

	AuthClaims struct {
		jwt.RegisteredClaims
		ID uint `json:"id"`
	}

	authClaimsKey struct{}

	middleware struct {
		jwtCfg config.JWT
	}
)

func NewMiddleware(jwtCfg config.JWT) Middleware {
	return middleware{jwtCfg}
}
