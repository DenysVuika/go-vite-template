package main

import (
	"fmt"
	"go-vite/api"
	"go-vite/frontend"
	"net/http"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	Version   = "dev" // Default value
	BuildDate = "unknown"
	GitCommit = "unknown"
)

const dbPath = "./data/go-vite.db"

func ApplyMigrations() {
	log.Info().Msg("Applying database migrations...")

	m, err := migrate.New(
		"file://data/migrations",
		"sqlite3://data/go--vite.db",
	)
	if err != nil {
		log.Fatal().Err(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err)
	}

	log.Info().Msg("Migrations applied successfully")
}

func main() {
	fmt.Printf("%s v.%s (Commit: %s, Built: %s)\n", "go-vite", Version, GitCommit, BuildDate)

	// Configure logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	// Load environment variables from .env file
	if _, err := os.Stat(".env"); err == nil {
		log.Info().Msg("Loading environment variables from .env file")
		godotenv.Load()
	} else if os.IsNotExist(err) {
		log.Info().Msg("No .env file found, using default environment variables")
	} else {
		log.Fatal().Err(err).Msg("Error checking for .env file")
	}

	// Apply database migrations
	ApplyMigrations()

	e := echo.New()

	// CORS configuration to allow all origins
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Allow all origins
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${remote_ip} - - [${time_rfc3339}] \"${method} ${uri} ${protocol}\" ${status} ${bytes_out}\n",
		Output: os.Stderr,
	}))
	e.Use(middleware.Recover())

	// Serve embedded static files
	frontend.RegisterFrontend(e)

	// API routes
	apiGroup := e.Group("/api")
	api.RegisterRoutes(apiGroup)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	log.Info().Msgf("Starting server on port %s", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
