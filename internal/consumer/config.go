package consumer

import (
	_ "embed"
	"log"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var configData []byte

type Config struct {
	KafkaHost      string `yaml:"kafka_host"`
	KafkaPort      int    `yaml:"kafka_port"`
	KafkaTopic     string `yaml:"kafka_topic"`
	CoreWebhookURL string `yaml:"core_webhook_url"`
}

var AppConfig Config

func (app *App) LoadConfig() {
	if app.hasError() {
		return
	}

	if err := yaml.Unmarshal(configData, &app.cfg); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}
}
