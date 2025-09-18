package core

import (
	_ "embed"
	"log/slog"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var configData []byte

type CoreConfig struct {
	EventServiceURL string `yaml:"event_service_url"`
}

var AppConfig CoreConfig

func (app *App) LoadConfig() {
	if app.hasError() {
		return
	}

	if err := yaml.Unmarshal(configData, &app.cfg); err != nil {
		slog.Error("Failed to parse core config.yaml", "error", err)
		// fallback to default
		app.cfg.EventServiceURL = "http://event:8081"
	}
}

func (app *App) GetEventServiceURL() string {
	if app.cfg.EventServiceURL != "" {
		return app.cfg.EventServiceURL
	}

	return "http://event:8081"
}
