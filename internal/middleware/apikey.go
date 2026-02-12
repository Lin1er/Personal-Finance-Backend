// Package middleware contains the HTTP middlewares for the API.
package middleware

import (
	"net/http"

	"personal-finance-backend/internal/service"
	"personal-finance-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

// APIKeyAuth returns a middleware that validates the X-API-Key header against the database.
func APIKeyAuth(apiKeyService *service.ApiKeyService) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")

		if key == "" {
			response.Error(c, http.StatusUnauthorized, "Missing API key")
			c.Abort()
			return
		}

		apiKey, err := apiKeyService.ValidateKey(c.Request.Context(), key)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid or inactive API key")
			c.Abort()
			return
		}

		// Store API key info in context for downstream handlers
		c.Set("api_key_id", apiKey.ID)
		c.Set("api_key_name", apiKey.Name)
		c.Next()
	}
}
