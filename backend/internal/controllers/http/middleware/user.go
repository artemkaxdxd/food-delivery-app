package middleware

import (
	"backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m middleware) HandleUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := m.extractToken(c)

		claims, err := m.parseToken(token, m.jwtCfg.Secret)
		if err != nil {
			m.handleError(c, http.StatusUnauthorized, config.CodeUnauthorized, err)
			return
		}

		c.Request = c.Request.WithContext(m.injectUserClaims(c.Request.Context(), AuthClaims{
			ID: claims.ID,
		}))
		c.Next()
	}
}
