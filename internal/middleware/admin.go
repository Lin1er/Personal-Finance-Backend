package middleware

import (
	"crypto/subtle"
	"net/http"

	"personal-finance-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

// AdminAuth validates the X-Admin-Key header against the ADMIN_API_KEY env variable.
// Used to protect API key management endpoints.
func AdminAuth(adminKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-Admin-Key")

		if key == "" {
			response.Error(c, http.StatusUnauthorized, "Missing admin key")
			c.Abort()
			return
		}

		if subtle.ConstantTimeCompare([]byte(key), []byte(adminKey)) != 1 {
			response.Error(c, http.StatusUnauthorized, "Invalid admin key")
			c.Abort()
			return
		}

		c.Next()
	}
}
