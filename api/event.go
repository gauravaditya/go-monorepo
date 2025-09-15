package api

type Event struct {
	Name      string `json:"name" yaml:"name"`
	Timestamp string `json:"timestamp" yaml:"timestamp"`
	Consumed  bool   `json:"consumed" yaml:"consumed"`
}
