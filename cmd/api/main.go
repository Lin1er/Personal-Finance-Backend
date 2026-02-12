// Package main implements the entry point for the Personal Finance Backend API server.
package main

import (
	"context"
	"log"

	"personal-finance-backend/internal/config"
	"personal-finance-backend/internal/db"
	"personal-finance-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.New(cfg)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer dbConn.Close()

	if err := dbConn.Ping(context.Background()); err != nil {
		log.Fatal("DB Error:", err)
	}

	log.Println("Successfully connected to the database!")

	r := gin.Default()

	handler.RegisterRoutes(r, dbConn, cfg)

	log.Println("Server running on :" + cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
