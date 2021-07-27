package gateway

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

type LineNotifyClient struct {
	BaseURL     string
	accessToken string
	HTTPClient  *http.Client
}

const (
	LineNotifyBaseURL = "https://notify-api.line.me"
)

func NewLineNotityClient() *LineNotifyClient {
	return &LineNotifyClient{
		BaseURL:     LineNotifyBaseURL,
		accessToken: "7qxDlbtL3M43QtTRHaFxQAFV3oWrPMGmJIZKuSxe357",
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *LineNotifyClient) SendMessage(message string) error {
	values := url.Values{}
	values.Set("message", message)

	req, err := http.NewRequest(
		"POST",
		c.BaseURL+"/api/notify/",
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		return nil
	}

	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, err = c.HTTPClient.Do(req)

	if err != nil {
		return err
	}
	return nil
}
