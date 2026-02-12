package handler

import (
	"net/http"

	"personal-finance-backend/internal/domain"
	"personal-finance-backend/internal/service"
	"personal-finance-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req domain.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create category")
		return
	}

	response.Success(c, http.StatusCreated, "Category created", category)
}

func (h *CategoryHandler) List(c *gin.Context) {
	categories, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to list categories")
		return
	}

	response.Success(c, http.StatusOK, "OK", categories)
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	category, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Category not found")
		return
	}

	response.Success(c, http.StatusOK, "OK", category)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var req domain.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.service.Update(c.Request.Context(), id, req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update category")
		return
	}

	response.Success(c, http.StatusOK, "Category updated", category)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete category")
		return
	}

	response.Success(c, http.StatusOK, "Category deleted", nil)
}
