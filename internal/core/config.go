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

func LoadConfig() {
	if err := yaml.Unmarshal(configData, &AppConfig); err != nil {
		slog.Error("Failed to parse core config.yaml", "error", err)
		// fallback to default
		AppConfig.EventServiceURL = "http://event:8081"
	}
}

func GetEventServiceURL() string {
	if AppConfig.EventServiceURL != "" {
		return AppConfig.EventServiceURL
	}
	return "http://event:8081"
}
