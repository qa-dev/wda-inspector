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

type Client struct {
	ip       string
	port     string
	bundleId string
}

func NewClient(ip string, port string, bundleId string) *Client {
	return &Client{ip: ip, port: port, bundleId: bundleId}
}

func (c *Client) BundleId() string {
	return c.bundleId
}

func (c *Client) url(uri string) string {
	return "http://" + c.ip + ":" + c.port + uri
}

func (c *Client) get(uri string) ([]byte, error) {
	r, err := http.Get(c.url(uri))
	return c.resonseProcess(r, err)
}

func (c *Client) post(uri string, data []byte) ([]byte, error) {
	r, err := http.Post(c.url(uri), "application/ajax", bytes.NewBuffer(data))
	return c.resonseProcess(r, err)
}

func (c *Client) resonseProcess(r *http.Response, err error) ([]byte, error) {
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
