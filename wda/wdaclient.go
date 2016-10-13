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

func NewClient(ip string, port string, bundleId string) *WdaClient {
	return &WdaClient{ip: ip, port: port, bundleId: bundleId}
}

func (c *WdaClient) BundleId() string {
	return c.bundleId
}

func (c *WdaClient) url(uri string) string {
	return "http://" + c.ip + ":" + c.port + uri
}

func (c *WdaClient) get(uri string) ([]byte, error) {
	r, err := http.Get(c.url(uri))
	return c.responseProcess(r, err)
}

func (c *WdaClient) post(uri string, data []byte) ([]byte, error) {
	r, err := http.Post(c.url(uri), "application/ajax", bytes.NewBuffer(data))
	return c.responseProcess(r, err)
}

func (c *WdaClient) responseProcess(r *http.Response, err error) ([]byte, error) {
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
