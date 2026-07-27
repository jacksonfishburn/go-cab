package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/jacksonfishburn/go-cab/internal/file"
)

type client struct {
	URL    string
	Token  string
	client *http.Client
}

func newClient(apiURL, token string) *client {
	return &client{
		URL:   apiURL,
		Token: token,
		client: &http.Client{
			Timeout: 5 * time.Minute,
		},
	}
}

func (c *client) request(method, path string, body []byte) ([]byte, int, error) {
	var reqBody io.Reader
	if body != nil {
		reqBody = bytes.NewReader(body)
	}

	endpoint := joinURL(c.URL, path)
	req, err := http.NewRequest(method, endpoint, reqBody)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	if body != nil {
		req.Header.Set("Content-Type", "application/octet-stream")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return respBody, resp.StatusCode, nil
}

func (c *client) Add(name string, blob []byte) (file.Record, error) {
	path := "add/" + url.PathEscape(name)

	respBody, status, err := c.request(http.MethodPost, path, blob)
	if err != nil {
		return file.Record{}, err
	}
	if err := checkStatus(status, respBody); err != nil {
		return file.Record{}, err
	}

	var record file.Record
	if err := json.Unmarshal(respBody, &record); err != nil {
		return file.Record{}, fmt.Errorf("decode response: %w", err)
	}
	return record, nil
}

func (c *client) Grab(name string) ([]byte, error) {
	path := "grab/" + url.PathEscape(name)

	respBody, status, err := c.request(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	if err := checkStatus(status, respBody); err != nil {
		return nil, err
	}
	return respBody, nil
}

func (c *client) Del(name string) error {
	path := "del/" + url.PathEscape(name)

	respBody, status, err := c.request(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	return checkStatus(status, respBody)
}

func (c *client) Peek() ([]file.Record, error) {
	respBody, status, err := c.request(http.MethodGet, "peek", nil)
	if err != nil {
		return nil, err
	}
	if err := checkStatus(status, respBody); err != nil {
		return nil, err
	}

	var list map[string]file.Record
	if err := json.Unmarshal(respBody, &list); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	records := make([]file.Record, 0, len(list))
	for _, r := range list {
		records = append(records, r)
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].Name < records[j].Name
	})
	return records, nil
}

func joinURL(base, path string) string {
	return strings.TrimRight(base, "/") + "/" + strings.TrimLeft(path, "/")
}

func checkStatus(status int, body []byte) error {
	if status >= 200 && status < 300 {
		return nil
	}

	var apiErr struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal(body, &apiErr); err == nil && apiErr.Error != "" {
		return fmt.Errorf("api error (%d): %s", status, apiErr.Error)
	}
	return fmt.Errorf("api error (%d): %s", status, bytes.TrimSpace(body))
}
