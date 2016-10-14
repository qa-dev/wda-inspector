package interfaces

type Requester interface {
	BundleId() string
	Url(uri string) string
	Get(uri string) ([]byte, error)
	Post(uri string, data []byte) ([]byte, error)
}
