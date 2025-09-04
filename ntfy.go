package ntfy

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

// Client represents an ntfy client
type Client struct {
	BaseURL string
	client  *http.Client
}

// New creates a new ntfy client
func New() *Client {
	baseURL := os.Getenv("NTFY_BASE_URL")

	if baseURL == "" {
		slog.Debug("NTFY_BASE_URL not set, using default https://ntfy.sh")
		baseURL = "https://ntfy.sh"
	}

	return &Client{
		BaseURL: baseURL,
		client:  &http.Client{},
	}
}

// MessageOptions contains options for sending a message
type MessageOptions struct {
	Message  string
	Title    string
	Priority int
	Tags     []string
	Click    string
	Attach   string
	Actions  string
	Email    string
	Icon     string
}

// PostWithOptions sends a message with additional options
func (c *Client) Post(topic string, opts MessageOptions) error {
	url := fmt.Sprintf("%s/%s", c.BaseURL, topic)
	log.Println(url)
	log.Println(filepath.Join(c.BaseURL, topic))

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(opts.Message))
	if err != nil {
		return err
	}

	// Set headers based on options
	if opts.Title != "" {
		req.Header.Set("Title", opts.Title)
	}
	if opts.Priority > 0 {
		req.Header.Set("Priority", fmt.Sprintf("%d", opts.Priority))
	}
	if len(opts.Tags) > 0 {
		tags := ""
		for i, tag := range opts.Tags {
			if i > 0 {
				tags += ","
			}
			tags += tag
		}
		req.Header.Set("Tags", tags)
	}
	if opts.Click != "" {
		req.Header.Set("Click", opts.Click)
	}
	if opts.Attach != "" {
		req.Header.Set("Attach", opts.Attach)
	}
	if opts.Actions != "" {
		req.Header.Set("Actions", opts.Actions)
	}
	if opts.Email != "" {
		req.Header.Set("Email", opts.Email)
	}
	if opts.Icon != "" {
		req.Header.Set("Icon", opts.Icon)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("failed to post message: status %d, body: %s", resp.StatusCode, body)
	}

	return nil
}
