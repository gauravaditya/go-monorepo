package consumer

import (
	"bytes"
	"log/slog"
	"net/http"
	"time"
)

// WebhookService is responsible for calling the core webhook

type WebhookService struct {
	Client *http.Client
}

func NewWebhookService() *WebhookService {
	return &WebhookService{
		Client: &http.Client{Timeout: 5 * time.Second},
	}
}

// Call posts the event payload to the core webhook URL
func (ws *WebhookService) Call(url string, payload []byte) error {
	slog.Info("WebhookService: posting to webhook", "url", url, "payload", string(payload))
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		slog.Error("WebhookService: failed to create request", "error", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := ws.Client.Do(req)
	if err != nil {
		slog.Error("WebhookService: HTTP request failed", "error", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		slog.Error("WebhookService: webhook returned non-200", "status", resp.StatusCode)
		return err
	}
	slog.Info("WebhookService: webhook call succeeded", "url", url)
	return nil
}
