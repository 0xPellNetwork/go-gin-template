// @title           Gin Template API
// @version         1.0
// @description     A clean and modern Go API template with Gin, GORM, and smart parameter binding
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

package main

import (
	"gin-template/pkg/config"
	"gin-template/pkg/database"
	"gin-template/pkg/router"

	"github.com/rs/zerolog/log"
)

func main() {
	// Initialize configuration
	cfg := config.New()

	// Setup logger
	config.SetupLogger(cfg.Log)
	logger := log.With().Str("component", "main").Logger()

	logger.Info().Msg("Starting Gin Template API Server")

	// Initialize database connection
	db, err := database.New(cfg.Database)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}
	logger.Info().Str("driver", cfg.Database.Driver).Msg("Database connected successfully")

	// Initialize router with new architecture
	r := router.New(db)

	// Start server
	logger.Info().
		Str("port", cfg.Server.Port).
		Str("log_level", cfg.Log.Level).
		Str("log_format", cfg.Log.Format).
		Msg("Server starting with new architecture")

	if err := r.Run(":" + cfg.Server.Port); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start server")
	}
}
