package wda

type Requester interface {
	BundleId() string
	Url(uri string) string
	Get(uri string) ([]byte, error)
	Post(uri string, data []byte) ([]byte, error)
}

type Client struct {
	Client Requester
}

func NewClient(r Requester) *Client {
	return &Client{Client: r}
}
