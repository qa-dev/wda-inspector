package wda

type Client struct {
	Client Requester
}

func NewBasicClient() *Client {
	return &Client{}
}

func (c *Client) SetClient(r Requester) {
	c.Client = r
}