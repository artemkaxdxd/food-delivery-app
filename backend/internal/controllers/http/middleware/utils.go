package middleware

import (
	"backend/config"
	"backend/internal/controllers/http/response"
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (m middleware) handleError(c *gin.Context, status int, code config.ServiceCode, err error) {
	c.AbortWithStatusJSON(status, response.NewErr(code, err.Error()))
}

func (m middleware) extractToken(c *gin.Context) string {
	token := c.Request.Header.Get("Authorization")
	strArr := strings.Split(token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (m middleware) parseToken(jwtToken, secret string) (*AuthClaims, error) {
	claims := &AuthClaims{}

	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, config.ErrInvalidToken
}

func (m middleware) injectUserClaims(ctx context.Context, c AuthClaims) context.Context {
	return context.WithValue(ctx, authClaimsKey{}, c)
}

func GetClaims(ctx context.Context) AuthClaims {
	claims, _ := ctx.Value(authClaimsKey{}).(AuthClaims)
	return claims
}
