package ntfy

// Message is a struct that represents a ntfy message
type Message struct { // TODO combine with server.message
	ID         string
	Event      string
	Time       int64
	Topic      string
	Message    string
	Title      string
	Priority   int
	Tags       []string
	Click      string
	Icon       string
	Attachment *Attachment

	// Additional fields
	TopicURL       string
	SubscriptionID string
	Raw            string
}

// Attachment represents a message attachment
type Attachment struct {
	Name    string `json:"name"`
	Type    string `json:"type,omitempty"`
	Size    int64  `json:"size,omitempty"`
	Expires int64  `json:"expires,omitempty"`
	URL     string `json:"url"`
	Owner   string `json:"-"` // IP address of uploader, used for rate limiting
}

type Priority = int

var PriorityMin = 1
var PriorityLow = 2
var PriorityDefault = 3
var PriorityHigh = 4
var PriorityMax = 5
