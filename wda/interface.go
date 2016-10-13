package wda

type Requester interface {
	BundleId() string
	url(uri string) string
	get(uri string) ([]byte, error)
	post(uri string, data []byte) ([]byte, error)
}
