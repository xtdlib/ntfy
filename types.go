package ntfy

type Message struct { // TODO combine with server.message
	Time     string `json:"time,omitempty"`
	Message  string `json:"message,omitempty"`
	Title    string `json:"title,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type Priority int
var PriorityMin = 1
var PriorityLow = 2
var PriorityDefault = 3
var PriorityHigh = 4
var PriorityMax = 5
