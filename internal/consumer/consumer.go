package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/gauravaditya/go-monorepo/api"
	"github.com/segmentio/kafka-go"
)

// StartConsumer starts the Kafka consumer loop and calls the webhook for each event.
func StartConsumer(ctx context.Context, webhookURL string, kafkaHost string, kafkaPort int) {
	broker := fmt.Sprintf("%s:%d", kafkaHost, kafkaPort)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		Topic:     AppConfig.KafkaTopic,
		Partition: 0,
		MinBytes:  1,    // 1B
		MaxBytes:  10e6, // 10MB
	})
	defer r.Close()

	webhookSvc := NewWebhookService()
	slog.Info("StartConsumer: started consuming", "broker", broker, "topic", AppConfig.KafkaTopic)
	attempt := 0
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			slog.Error("StartConsumer: error reading message", "error", err)
			time.Sleep(backoff(attempt))
			attempt++
			continue
		}
		var event api.Event
		if err := json.Unmarshal(m.Value, &event); err != nil {
			slog.Error("StartConsumer: failed to unmarshal event from Kafka", "error", err, "raw", string(m.Value))
			continue
		}
		slog.Info("StartConsumer: event received", "event", event)
		event.Consumed = true // mark event as consumed
		payload, err := json.Marshal(event)
		if err != nil {
			slog.Error("StartConsumer: failed to marshal event for webhook", "error", err)
			continue
		}
		if err := webhookSvc.Call(webhookURL, payload); err != nil {
			slog.Error("StartConsumer: webhook call failed", "error", err, "webhook", webhookURL)
			continue
		}
		slog.Info("StartConsumer: webhook call succeeded", "webhook", webhookURL)
	}
}

func backoff(attempt int) time.Duration {
	baseDelay := 1 * time.Second
	maxDelay := 10 * time.Second
	delay := time.Duration(attempt) * baseDelay
	if delay > maxDelay {
		delay = maxDelay
	}
	return delay
}
