package main

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

type Client struct {
	URL    string
	Token  string
	client *http.Client
}

func NewClient(URL, token string) *Client {
	return &Client{
		URL:   URL,
		Token: token,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Request(method, path string, body []byte) ([]byte, int, error) {
	var reqBody io.Reader
	if body != nil {
		reqBody = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, c.URL+path, reqBody)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

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

type Record struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Size      int       `json:"size"`
	MD5       string    `json:"md5"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Client) Add(name string, blob []byte) (Record, error) {
	return Record{}, nil
}
