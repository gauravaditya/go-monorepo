package event

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var configData []byte

type Config struct {
	KafkaHost  string `yaml:"kafka_host"`
	KafkaPort  int    `yaml:"kafka_port"`
	KafkaTopic string `yaml:"kafka_topic"`
}

var AppConfig Config

func (app *App) LoadConfig() {
	if app.hasError() {
		return
	}

	if err := yaml.Unmarshal(configData, &app.cfg); err != nil {
		app.err = fmt.Errorf("failed to parse event config.yaml: %w", err)
	}
}
