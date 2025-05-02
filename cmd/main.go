package main

import (
	_ "SentinelAuth/docs"
	"SentinelAuth/internal/config"
	"SentinelAuth/internal/handler"
	"SentinelAuth/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load .env variables
	config.LoadEnv()

	// Init PostgreSQL and Redis
	db := infrastructure.InitPostgres()
	rdb := infrastructure.InitRedis()
	infrastructure.MigrateDB()

	// Router
	router := gin.Default()

	// Swagger
	handler.RegisterSwaggerRoutes(router)

	// Correct usage: pass all required arguments
	handler.RegisterRoutes(router, db, rdb)

	// Run server
	router.Run(":" + config.GetEnv("PORT", "8080"))
}
