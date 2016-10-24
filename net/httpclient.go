package net

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	StatusOK = 0
)

type Requester interface {
	Url(uri string) string
	Get(uri string) ([]byte, error)
	Post(uri string, data []byte) ([]byte, error)
}

type HttpClient struct {
	ip       string
	port     string
}

func NewHttpClient(ip string, port string) *HttpClient {
	return &HttpClient{ip: ip, port: port}
}


func (c *HttpClient) Url(uri string) string {
	return "http://" + c.ip + ":" + c.port + uri
}

func (c *HttpClient) Get(uri string) ([]byte, error) {
	r, err := http.Get(c.Url(uri))
	return c.processResponse(r, err)
}

func (c *HttpClient) Post(uri string, data []byte) ([]byte, error) {
	r, err := http.Post(c.Url(uri), "application/ajax", bytes.NewBuffer(data))
	return c.processResponse(r, err)
}

func (c *HttpClient) processResponse(r *http.Response, err error) ([]byte, error) {
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
