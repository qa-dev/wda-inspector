package wda

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type FakeGettextRequesterClient struct {}

func (c *FakeGettextRequesterClient) BundleId() string {
	return "fake.bundleid"
}

func (c *FakeGettextRequesterClient) url(uri string) string {
	fmt.Println("http://fakeurl.fake:80/" + uri)
	return "http://fakeurl.fake:80/" + uri
}

func (c *FakeGettextRequesterClient) get(uri string) ([]byte, error) {
	var res string

	switch uri {
	case "/status":
		res = "{\"sessionId\": \"fakesession\", \"status\": 0}"
	case "/session/fakesession/element/fakeelementid/text":
		res = "{\"value\": \"FakeText\", \"status\": 0}"
	default:
		res = "{\"error\": 1}"
	}
	return []byte(res), nil
}

func (c *FakeGettextRequesterClient) post(uri string, data []byte) ([]byte, error) {
	var res string

	switch uri {
	case "/session/fakesession/element":
		res = "{\"value\": {\"ELEMENT\": \"fakeelementid\", \"type\": \"faketype\"}, \"status\": 0}"
	default:
		res = "{\"error\": 1}"
	}

	return []byte(res), nil
}

func TestClient_GetText(t *testing.T) {
	requester := &FakeGettextRequesterClient{}
	client := NewBasicClient()
	client.SetClient(requester)
	findResp, err := client.Find("xui", "findme")
	assert.NoError(t, err, "Error finding fake element")
	elementId := findResp.Value.ElementId
	textResp, err := client.GetText(elementId)
	assert.NoError(t, err, "Error getting fake element text")
	assert.Equal(t, "FakeText", textResp.Value, "Text not equals")
}
