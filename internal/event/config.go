package event

import (
	_ "embed"
	"log"

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

func LoadConfig() {
	if err := yaml.Unmarshal(configData, &AppConfig); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}
}
