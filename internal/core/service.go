package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type RegisterEventsRequest struct {
	Count int `json:"count"`
}

type EventServiceResponse struct {
	Message string `json:"message"`
}

// CallEventService publishes events by calling the event service REST API
func CallEventService(eventServiceURL string, count int) error {
	payload := RegisterEventsRequest{Count: count}
	body, err := json.Marshal(payload)
	if err != nil {
		slog.Error("CallEventService: failed to marshal payload", "error", err)
		return err
	}
	resp, err := http.Post(fmt.Sprintf("%s/produce", eventServiceURL), "application/json", bytes.NewReader(body))
	if err != nil {
		slog.Error("CallEventService: failed to call event service", "error", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		slog.Error("CallEventService: event service returned non-200", "status", resp.StatusCode)
		return fmt.Errorf("event service error: %s", resp.Status)
	}
	slog.Info("CallEventService: events published via event service", "count", count)
	return nil
}
