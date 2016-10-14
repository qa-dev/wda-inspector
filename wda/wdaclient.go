package wda

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	StatusOK = 0
)

type WdaClient struct {
	ip       string
	port     string
	bundleId string
}

func NewWdaClient(ip string, port string, bundleId string) *WdaClient {
	return &WdaClient{ip: ip, port: port, bundleId: bundleId}
}

func (c *WdaClient) BundleId() string {
	return c.bundleId
}

func (c *WdaClient) Url(uri string) string {
	return "http://" + c.ip + ":" + c.port + uri
}

func (c *WdaClient) Get(uri string) ([]byte, error) {
	r, err := http.Get(c.Url(uri))
	return c.processResponse(r, err)
}

func (c *WdaClient) Post(uri string, data []byte) ([]byte, error) {
	r, err := http.Post(c.Url(uri), "application/ajax", bytes.NewBuffer(data))
	return c.processResponse(r, err)
}

func (c *WdaClient) processResponse(r *http.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return nil, errors.New("Bad status from inspector: " + r.Status)
	}
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return res, nil
}
