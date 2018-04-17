// Package webhook provides a HTTP/POST webhook implementation.

package httplog

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
)

// Client is a slack webhook client for posting messages using a webhook URL.
type Client struct {
	// URL is the webhook URL to use
	URL string
}

// New returns a new Client which sends request using the webhook URL.
func New(url string) *Client {
	return &Client{URL: url}
}

// Send sends the request to slack using the webhook protocol.
// The url parameter only exists to satisfy the slack.Client interface
// and is not used by the webhook Client.
func (c *Client) Send(msg, resp interface{}) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	r, err := http.Post(c.URL, "application/json; charset=utf-8", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(r.Body)
		return fmt.Errorf(string(b))
	}

	return nil
}