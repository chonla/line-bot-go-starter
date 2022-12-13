package models

type Event struct {
	ReplyToken      string           `json:"replyToken"`
	Type            string           `json:"type"`
	Mode            string           `json:"mode"`
	Timestamp       uint             `json:"timestamp"`
	Source          *EventSource     `json:"source"`
	WebhookEventID  string           `json:"webhookEventId"`
	DeliveryContext *DeliveryContext `json:"deliveryContext,omitempty"`
	Beacon          *Beacon          `json:"beacon,omitempty"`
}
