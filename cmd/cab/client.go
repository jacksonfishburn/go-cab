package main

import (
	"net/http"
	"time"
)

type Client struct {
	URL    string
	Token     string
	HTTPClient *http.Client
}

func NewClient(URL, token string) *Client {
	return &Client{
		URL: URL,
		Token:  token,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type Record struct {
	Name      string    `json:"name"`
	Size      int       `json:"size"`
	MD5       string    `json:"md5"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func (c *Client) Add(name string, blob []byte) (Record, error) {
	return Record{}, nil	
}