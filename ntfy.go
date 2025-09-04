package ntfy

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

// Client represents an ntfy client
type Client struct {
	BaseURL    string
	HttpClient *http.Client
}

// New creates a new ntfy client
func New(baseURL string, httpClient *http.Client) *Client {
	if baseURL == "" {
		slog.Debug("using default https://ntfy.sh")
		baseURL = "https://ntfy.sh"
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		BaseURL:    baseURL,
		HttpClient: httpClient,
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

var DefaultClient *Client

func init() {
	baseURL := os.Getenv("NTFY_BASE_URL")
	if baseURL == "" {
		baseURL = "https://ntfy.sh"
	}

	DefaultClient = New(baseURL, http.DefaultClient)
}

func Post(topic string, opts MessageOptions) error {
	return DefaultClient.Post(topic, opts)
}

// PostWithOptions sends a message with additional options
func (c *Client) Post(topic string, opts MessageOptions) error {
	url := fmt.Sprintf("%s/%s", c.BaseURL, topic)

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

	resp, err := c.HttpClient.Do(req)
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
