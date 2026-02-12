// Package handler contains the HTTP handlers for the API.package handler
package handler

import (
	"net/http"

	"personal-finance-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(c *gin.Context) {
	response.Success(c, http.StatusOK, "ok", nil)
}

func (h *HealthHandler) Ping(c *gin.Context) {
	response.Success(c, http.StatusOK, "pong", nil)
}
