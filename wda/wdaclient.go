package wda

type Requester interface {
	Url(uri string) string
	Get(uri string) ([]byte, error)
	Post(uri string, data []byte) ([]byte, error)
}

type Client struct {
	Client   Requester
	bundleId string
}

func NewClient(r Requester, bundleId string) *Client {
	return &Client{Client: r, bundleId: bundleId}
}

func (c *Client) BundleId() string {
	return c.bundleId
}
