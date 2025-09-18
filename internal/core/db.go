package core

import (
	"fmt"
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func (app *App) InitDB() {
	if app.hasError() {
		return
	}

	dsn := os.Getenv("CORE_POSTGRES_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=eventsdb port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		app.err = fmt.Errorf("failed to connect to Postgres: %w", err)

		return
	}

	if err := db.AutoMigrate(&Event{}); err != nil {
		app.err = fmt.Errorf("failed to migrate Event table: %w", err)

		return
	}

	app.db = db
	slog.Info("Connected to Postgres and migrated Event table")
}
