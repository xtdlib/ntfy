package ntfy_test

import (
	"testing"

	"github.com/xtdlib/ntfy"
)

func TestClient(t *testing.T) {
	// Send a complex message with options
	err := ntfy.Post("system", ntfy.MessageOptions{
		Message: `There's someone at the door. 🐶

Please check if it's a good boy or a hooman. 
Doggies have been known to ring the doorbell.`,
		// Click:   "https://home.nest.com/",
		// Attach:  "https://nest.com/view/yAxkasd.jpg",
		// Actions: "http, Open door, https://api.nest.com/open/yAxkasd, clear=true",
		// Email:   "phil@example.com",
	})
	if err != nil {
		panic(err)
	}
}
