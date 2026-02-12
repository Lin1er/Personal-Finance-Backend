package handler

import (
	"net/http"

	"personal-finance-backend/internal/domain"
	"personal-finance-backend/internal/service"
	"personal-finance-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type ApiKeyHandler struct {
	service *service.ApiKeyService
}

func NewApiKeyHandler(s *service.ApiKeyService) *ApiKeyHandler {
	return &ApiKeyHandler{service: s}
}

// Create godoc
// POST /api/v1/api-keys
// Body: { "name": "my-frontend-app" }
// Response: returns the full key (only time it's shown)
func (h *ApiKeyHandler) Create(c *gin.Context) {
	var req domain.CreateApiKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	apiKey, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create API key")
		return
	}

	response.Success(c, http.StatusCreated, "API key created. Save this key, it won't be shown again.", apiKey)
}

// List godoc
// GET /api/v1/api-keys
func (h *ApiKeyHandler) List(c *gin.Context) {
	keys, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to list API keys")
		return
	}

	response.Success(c, http.StatusOK, "OK", keys)
}

// GetByID godoc
// GET /api/v1/api-keys/:id
func (h *ApiKeyHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	apiKey, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "API key not found")
		return
	}

	response.Success(c, http.StatusOK, "OK", apiKey)
}

// Update godoc
// PATCH /api/v1/api-keys/:id
// Body: { "name": "new-name", "is_active": false }
func (h *ApiKeyHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var req domain.UpdateApiKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	apiKey, err := h.service.Update(c.Request.Context(), id, req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update API key")
		return
	}

	response.Success(c, http.StatusOK, "API key updated", apiKey)
}

// Delete godoc
// DELETE /api/v1/api-keys/:id
func (h *ApiKeyHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete API key")
		return
	}

	response.Success(c, http.StatusOK, "API key deleted", nil)
}
