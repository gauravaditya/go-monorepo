package core

import (
	"log/slog"

	"github.com/gauravaditya/go-monorepo/api"
)

type Event struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"index"`
	Timestamp string
	Consumed  bool
}

// CreateEvent inserts a new event into the database
func CreateEvent(name, timestamp string) error {
	event := Event{Name: name, Timestamp: timestamp, Consumed: false}
	if err := DB.Create(&event).Error; err != nil {
		slog.Error("CreateEvent: failed to insert event", "error", err, "name", name, "timestamp", timestamp)
		return err
	}
	slog.Info("CreateEvent: event inserted", "name", name, "timestamp", timestamp)
	return nil
}

// UpdateEventConsumed updates the consumed count for an event
func UpdateEventConsumed(name, timestamp string, consumed bool) error {
	res := DB.Model(&Event{}).Where("name = ? AND timestamp = ?", name, timestamp).Update("consumed", consumed)
	if res.Error != nil {
		slog.Error("UpdateEventConsumed: failed to update event", "error", res.Error, "name", name, "timestamp", timestamp)
		return res.Error
	}
	if res.RowsAffected == 0 {
		slog.Warn("UpdateEventConsumed: no event found to update", "name", name, "timestamp", timestamp)
	} else {
		slog.Info("UpdateEventConsumed: event updated", "name", name, "timestamp", timestamp, "consumed", consumed)
	}
	return nil
}

// GetAllEvents fetches all events from the database
func GetAllEvents() ([]api.Event, error) {
	var dbEvents []Event
	err := DB.Find(&dbEvents).Error
	if err != nil {
		slog.Error("GetAllEvents: failed to fetch events", "error", err)
		return nil, err
	}
	var result []api.Event
	for _, e := range dbEvents {
		result = append(result, api.Event{
			Name:      e.Name,
			Timestamp: e.Timestamp,
			Consumed:  e.Consumed,
		})
	}
	slog.Info("GetAllEvents: fetched events", "count", len(result))
	return result, nil
}
