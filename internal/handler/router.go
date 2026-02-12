package handler

import (
	"personal-finance-backend/internal/config"
	"personal-finance-backend/internal/middleware"
	"personal-finance-backend/internal/repository"
	"personal-finance-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// RegisterRoutes registers all API routes.
func RegisterRoutes(r *gin.Engine, db *pgxpool.Pool, cfg *config.Config) {
	// Init layers
	// =========================
	// API Key management
	// ==========================
	apiKeyRepo := repository.NewApiKeyRepository(db)
	apiKeyService := service.NewApiKeyService(apiKeyRepo)
	apiKeyHandler := NewApiKeyHandler(apiKeyService)
	// =========================
	// Categories
	// ==========================
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)
	// =========================
	// Transactions
	// ==========================
	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := NewTransactionHandler(transactionService)
	// =========================
	// Health check
	// ==========================
	healthHandler := NewHealthHandler()

	// Public routes
	r.GET("/health", healthHandler.Check)

	// Admin routes (protected by ADMIN_API_KEY from env)
	// Used to manage API keys (create, list, update, delete)
	admin := r.Group("/admin/v1")
	admin.Use(middleware.AdminAuth(cfg.AdminAPIKey))
	{
		admin.POST("/api-keys", apiKeyHandler.Create)
		admin.GET("/api-keys", apiKeyHandler.List)
		admin.GET("/api-keys/:id", apiKeyHandler.GetByID)
		admin.PATCH("/api-keys/:id", apiKeyHandler.Update)
		admin.DELETE("/api-keys/:id", apiKeyHandler.Delete)
	}

	// API routes (protected by API key from database)
	// Used by external projects/services
	api := r.Group("/api/v1")
	api.Use(middleware.APIKeyAuth(apiKeyService))
	{
		api.GET("/ping", healthHandler.Ping)

		// Categories
		api.POST("/categories", categoryHandler.Create)
		api.GET("/categories", categoryHandler.List)
		api.GET("/categories/:id", categoryHandler.GetByID)
		api.PATCH("/categories/:id", categoryHandler.Update)
		api.DELETE("/categories/:id", categoryHandler.Delete)

		// Transactions
		api.POST("/transactions", transactionHandler.Create)
		api.GET("/transactions", transactionHandler.List)
		api.GET("/transactions/:id", transactionHandler.GetByID)
		api.PATCH("/transactions/:id", transactionHandler.Update)
		api.DELETE("/transactions/:id", transactionHandler.Delete)
	}
}
