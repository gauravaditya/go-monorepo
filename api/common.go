package api

type Event struct {
	Name      string `json:"name" yaml:"name"`
	Timestamp string `json:"timestamp" yaml:"timestamp"`
	Consumed  bool   `json:"consumed" yaml:"consumed"`
}

type RegisterRequest struct {
	Count int `json:"count"`
}

// RegisterResponse is the response for /register endpoint
type RegisterResponse struct {
	Message string  `json:"message"`
	Events  []Event `json:"events"`
}

// EventsDataResponse is the response for /events-data endpoint
type EventsDataResponse struct {
	Events []Event `json:"events"`
}
