package wda

import "github.com/qa-dev/wda-inspector/interfaces"

type Client struct {
	Client interfaces.Requester
}

func NewClient(r interfaces.Requester) *Client {
	return &Client{Client: r}
}
