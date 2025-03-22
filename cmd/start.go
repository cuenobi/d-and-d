package cmd

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"d-and-d/config"
	_ "d-and-d/docs"
	"d-and-d/internal/adapter/jwt"
	"d-and-d/internal/adapter/postgres"
	"d-and-d/internal/routes"
	"d-and-d/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/spf13/cobra"
)

func start(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	cfg := config.GetConfig()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: cfg.Log.Level,
	}))

	logger.Info("App is starting...")

	// Init database
	pg := postgres.NewPostgres(cfg.Postgres, ctx, logger)
	postgres.SeedDatabase(pg, logger)

	// Init fiber server
	f := startServer(cfg.ServerConfig.Port, logger, cfg)

	// Init repositories
	userRepo := postgres.NewUser(pg)
	publicRepo := postgres.NewPublic(pg)
	optionRepo := postgres.NewOption(pg)

	// Init services
	jwtService := jwt.NewJwtToken(cfg.JwtConfig)
	userService := service.NewUserService(userRepo, optionRepo, jwtService, logger, cfg)
	publicService := service.NewPublicService(publicRepo)

	validate := validator.New()

	// Register routes
	routes.NewRouteUserHandler(f, userService, jwtService, validate, logger)
	routes.NewRoutePublicHandler(f, publicService, logger, cfg)

	gracefulShutdown(f, logger)

	return nil
}

// @title RESTful API for a Dungeons & Dragons (D&D)
// @version 1.0
// @description Character and quest management application using the Go language and any suitable framework
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @security BearerAuth
func startServer(port string, slog *slog.Logger, cfg *config.Config) *fiber.App {
	// Create a new Fiber application
	f := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	f.Use(cors.New())

	f.Use(recover.New(recover.Config{
		EnableStackTrace: cfg.ServerConfig.EnableStackTrace,
	}))

	loggerMiddleware := logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
		Format:     "${time} | ${status} | ${latency} | ${ips} | ${method} | ${path}\n",
	})
	f.Use(loggerMiddleware)

	// Define a simple ping route for health checks
	f.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Swagger route
	f.Get("/swagger/*", swagger.HandlerDefault)

	// Start the server in a goroutine
	go func() {
		// Listen on the specified port
		if err := f.Listen(":" + port); err != nil {
			// Log any errors encountered while starting the server
			slog.With("error", err).Error("Error starting server")
		}
	}()

	// Return the Fiber application instance
	return f
}

func gracefulShutdown(f *fiber.App, logger *slog.Logger) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sigReceived := <-stop
	logger.Info("Received signal", slog.String("signal", sigReceived.String()))

	// The shutdown deadline is set to 3 seconds. If the server does not
	// shut down within this deadline, the program will exit with code 0.
	shutdownDeadline := time.After(3 * time.Second)

	// Shut down the server gracefully.
	if err := f.Shutdown(); err != nil {
		logger.Error("Error shutting down server", slog.Any("error", err))
	}

	// Wait until the server is fully shut down or the deadline is reached.
	select {
	case <-shutdownDeadline:
		logger.Info("Graceful shutdown completed")
	case <-time.After(5 * time.Second):
		logger.Info("Graceful shutdown timed out")
	}

	// Exit the program with code 0.
	os.Exit(0)
}
