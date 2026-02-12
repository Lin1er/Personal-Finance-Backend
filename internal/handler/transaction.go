package handler

import (
	"net/http"

	"personal-finance-backend/internal/domain"
	"personal-finance-backend/internal/service"
	"personal-finance-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

func (h *TransactionHandler) Create(c *gin.Context) {
	var req domain.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	tx, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create transaction: "+err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Transaction created", tx)
}

func (h *TransactionHandler) List(c *gin.Context) {
	var filter domain.TransactionFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	transactions, total, err := h.service.GetAll(c.Request.Context(), filter)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to list transactions")
		return
	}

	response.Success(c, http.StatusOK, "OK", gin.H{
		"transactions": transactions,
		"total":        total,
		"page":         filter.Page,
		"limit":        filter.Limit,
	})
}

func (h *TransactionHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	tx, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Transaction not found")
		return
	}

	response.Success(c, http.StatusOK, "OK", tx)
}

func (h *TransactionHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var req domain.UpdateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	tx, err := h.service.Update(c.Request.Context(), id, req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update transaction")
		return
	}

	response.Success(c, http.StatusOK, "Transaction updated", tx)
}

func (h *TransactionHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete transaction")
		return
	}

	response.Success(c, http.StatusOK, "Transaction deleted", nil)
}
