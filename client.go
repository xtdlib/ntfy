package ntfy

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

func send(topic string, msg Message) {
	body, err := json.Marshal(msg)
	if err != nil {
		slog.Error("ntfy: failed to marshal message", "error", err)
		return
	}
	resp, err := http.DefaultClient.Post("https://ping.xkor.stream/"+topic, "", bytes.NewReader(body))
	if err != nil {
		slog.Error("ntfy: failed to send message", "error", err)
		return
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		slog.Error("ntfy: unexpected status", "status", resp.StatusCode)
	}
}

func buildMessage(msg string, priority ...Priority) Message {
	m := Message{Message: msg}
	if len(priority) != 0 {
		m.Priority = int(priority[0])
	}
	now := time.Now()
	dayOfWeek := map[time.Weekday]string{
		time.Sunday: "일", time.Monday: "월", time.Tuesday: "화",
		time.Wednesday: "수", time.Thursday: "목", time.Friday: "금",
		time.Saturday: "토",
	}[now.Weekday()]
	m.Time = dayOfWeek + " " + now.Format(time.DateTime) + "\n"
	return m
}

func Send(topic string, msg string, priority ...Priority) {
	go send(topic, buildMessage(msg, priority...))
}

func SendSync(topic string, msg string, priority ...Priority) {
	send(topic, buildMessage(msg, priority...))
}
