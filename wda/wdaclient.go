package wda

import "github.com/qa-dev/wda-inspector/net"

type Client struct {
	httpClient net.Requester
	bundleId   string
}

func NewClient(r net.Requester, bundleId string) *Client {
	return &Client{httpClient: r, bundleId: bundleId}
}

func (c *Client) BundleId() string {
	return c.bundleId
}
