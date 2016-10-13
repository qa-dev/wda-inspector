package wda

type Client struct {
	Client Requester
}

func NewClient(r Requester) *Client {
	return &Client{Client: r}
}
