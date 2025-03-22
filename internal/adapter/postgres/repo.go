package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"d-and-d/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Host         string
	Port         string
	Name         string
	Username     string
	Password     string
	SeedPassword string
}

func NewPostgres(cfg *PostgresConfig, ctx context.Context, logger *slog.Logger) *gorm.DB {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)

	var db *gorm.DB
	var err error

	// Retry up to 5 times
	for i := 0; i < 5; i++ {
		// Open the connection
		db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err == nil {
			// Successfully connected, proceed with migration
			break
		}

		// If connection fails, log the error and wait before retrying
		logger.Info("Failed to connect to DB", "attempt", i+1, "error", err)
		time.Sleep(2 * time.Second) // Wait for 2 seconds before retrying
	}

	if err != nil {
		logger.Warn("Could not connect to DB after 5 attempts", "error", err)
	}

	// Perform database migrations
	if err := db.AutoMigrate(
		model.Character{},
		model.Class{},
		model.DifficultyLevels{},
		model.Quest{},
		model.Race{},
		model.User{},
		model.UserQuest{},
	); err != nil {
		logger.Error("Migration failed: ", "error", err)
	}

	return db
}

func SeedDatabase(db *gorm.DB, logger *slog.Logger) {
	seedFilePath := "/root/seed.sql"
	data, err := os.ReadFile(seedFilePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read seed file: %v", err))
	}

	sql := string(data)

	if err := db.Exec(sql).Error; err != nil {
		logger.Warn("Failed to seed database", "error", err)
		return
	}

	logger.Info("Database seeding completed!")
}
