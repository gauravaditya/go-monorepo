package core

import (
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := os.Getenv("CORE_POSTGRES_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=eventsdb port=5432 sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to Postgres", "error", err)
		return err
	}
	if err := db.AutoMigrate(&Event{}); err != nil {
		slog.Error("Failed to migrate Event table", "error", err)
		return err
	}
	DB = db
	slog.Info("Connected to Postgres and migrated Event table")
	return nil
}
