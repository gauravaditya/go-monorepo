package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gauravaditya/go-monorepo/api"
)

type RegisterEventsRequest struct {
	Count int `json:"count"`
}

type EventServiceResponse struct {
	Message string `json:"message"`
}

// CallEventServiceWithPayload sends a single api.Event to the event service /produce endpoint
func (app *App) CallEventServiceWithPayload(eventServiceURL string, ev api.Event) error {
	body, err := json.Marshal(ev)
	if err != nil {
		slog.Error("CallEventServiceWithPayload: failed to marshal event", "error", err)
		return err
	}
	resp, err := http.Post(fmt.Sprintf("%s/produce", eventServiceURL), "application/json", bytes.NewReader(body))
	if err != nil {
		slog.Error("CallEventServiceWithPayload: failed to call event service", "error", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		slog.Error("CallEventServiceWithPayload: event service returned non-200", "status", resp.StatusCode)
		return fmt.Errorf("event service error: %s", resp.Status)
	}
	slog.Info("CallEventServiceWithPayload: event sent to event service", "event", ev)
	return nil
}
