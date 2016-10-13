package wda

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type FakeGettextRequesterClient struct{}

func (c *FakeGettextRequesterClient) BundleId() string {
	return "fake.bundleid"
}

func (c *FakeGettextRequesterClient) url(uri string) string {
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
		return nil, errors.New("Error doing request")
	}
	return []byte(res), nil
}

func (c *FakeGettextRequesterClient) post(uri string, data []byte) ([]byte, error) {
	var res string

	switch uri {
	case "/session/fakesession/element":
		var findResp FindRequest
		err := json.Unmarshal(data, &findResp)
		if err == nil && findResp.Value == "findme" {
			res = "{\"value\": {\"ELEMENT\": \"fakeelementid\", \"type\": \"faketype\"}, \"status\": 0}"
		} else {
			return nil, errors.New("Error doing request")
		}
	default:
		return nil, errors.New("Error doing request")
	}
	return []byte(res), nil
}

func TestClient_GetText(t *testing.T) {
	requester := &FakeGettextRequesterClient{}
	client := NewBasicClient()
	client.SetClient(requester)

	// test success
	findResp, err := client.Find("xui", "findme")
	assert.NoError(t, err, "Error finding fake element")
	elementId := findResp.Value.ElementId
	textResp, err := client.GetText(elementId)
	assert.NoError(t, err, "Error getting fake element text")
	assert.Equal(t, "FakeText", textResp.Value, "Text not equals")

	// test error
	_, err = client.Find("xui", "not_exists_element")
	assert.Error(t, err, "Element must not be found")
	_, err = client.GetText("no_element")
	assert.Error(t, err, "Element text must not be found")
}
