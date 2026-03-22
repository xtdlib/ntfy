package ntfy

type Message struct { // TODO combine with server.message
	Time       string
	Topic      string
	Message    string
	Title      string
	Priority   int
}

var PriorityMin = 1
var PriorityLow = 2
var PriorityDefault = 3
var PriorityHigh = 4
var PriorityMax = 5
