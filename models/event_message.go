package models

type EventMessage struct {
	Destination string  `json:"destination"`
	Events      []Event `json:"events"`
}
