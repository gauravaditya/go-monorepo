package event

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

// ProduceEventMessage publishes a single event JSON payload to Kafka
func ProduceEventMessage(payload []byte, eventName string) error {
	broker := fmt.Sprintf("%s:%d", AppConfig.KafkaHost, AppConfig.KafkaPort)
	w := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    AppConfig.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
	defer w.Close()

	msg := kafka.Message{
		Key:   []byte(eventName),
		Value: payload,
	}
	if err := w.WriteMessages(context.Background(), msg); err != nil {
		slog.Error("ProduceEventMessage: failed to write message", "error", err, "event", eventName)
		return err
	}
	slog.Info("ProduceEventMessage: event produced", "event", eventName)
	return nil
}
