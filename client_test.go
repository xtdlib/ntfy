package ntfy

import (
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	Send("xaut.low", "Hello world!")
	time.Sleep(time.Second)
}
