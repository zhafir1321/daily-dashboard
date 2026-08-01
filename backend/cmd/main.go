package main

import (
	serve "backend/api/server"
	routes "backend/api/server/router"
	"backend/configs"
	"backend/internal/app/database"
	"backend/internal/app/handlers"
	"backend/internal/app/repositories"
	"backend/internal/app/services"
	"backend/internal/middleware"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	config := configs.NewConfig()

	client, err := database.NewSQLClient(database.Config{
		DBDriver:          config.Database.DatabaseDriver,
		DBSource:          config.Database.DatabaseSource,
		MaxOpenConns:      25,
		MaxIdleConns:      25,
		ConnMaxIdleTime:   15 * time.Minute,
		ConnectionTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database client")
		return
	}

	defer func() {
		if err := client.Close(); err != nil {
			log.Error().Msgf("Failed to close database client: %v", err)
		}
	}()

	if err := database.RunMigrations(config.Database.DatabaseSource); err != nil {
		log.Fatal().Err(err).Msg("Failed to run database migrations")
		return
	}

	// Initialize logger
	baseLogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	todoServiceLogger := baseLogger.With().Str("component", "Todo Service").Logger()
	todoHandlerLogger := baseLogger.With().Str("component", "Todo Handler").Logger()
	eventHandlerLogger := baseLogger.With().Str("component", "Event Handler").Logger()

	// Initialize repositories
	userRepository := repositories.NewUserRepository(client.DB)
	todoRepository := repositories.NewTodoRepository(client.DB)
	eventRepository := repositories.NewEventRepository(client.DB)

	// Initialize middleware
	middleware := middleware.NewMiddleware(userRepository, config.JWT.JWTSecret)

	// Initialize services
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(userRepository, config.JWT.JWTSecret)
	todoService := services.NewTodoService(todoRepository, todoServiceLogger)
	eventService := services.NewEventService(eventRepository, baseLogger)

	// Pass services to handlers
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)
	todoHandler := handlers.NewTodoHandler(todoService, todoHandlerLogger)
	eventHandler := handlers.NewEventHandler(eventService, eventHandlerLogger)

	cors := config.CorsNew()

	router := gin.Default()
	router.Use(cors)

	// Register routes
	routes.RegisterAuthEndpoints(router, authHandler)

	routes.RegisterPublicEndpoints(router, userHandler, middleware)
	routes.RegisterTodoEndpoints(router, todoHandler, middleware)
	routes.RegisterEventEndpoints(router, eventHandler, middleware)

	server := serve.NewServer(log.Logger, router, config)
	server.Serve()

}
