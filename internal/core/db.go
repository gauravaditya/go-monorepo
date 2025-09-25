package core

import (
	"fmt"
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func thisOrDefault(key, defaultVal string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return defaultVal
}

func (app *App) InitDB() {
	if app.hasError() {
		return
	}

	host := thisOrDefault("DB_HOST", "localhost")
	port := thisOrDefault("DB_PORT", "5432")
	usr := thisOrDefault("DB_USER", "postgres")
	pwd := thisOrDefault("DB_HOST", "postgres")
	dbname := thisOrDefault("DB_NAME", "postgres")
	sslmode := thisOrDefault("SSL_MODE", "disable")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, usr, pwd, dbname, port, sslmode)

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
