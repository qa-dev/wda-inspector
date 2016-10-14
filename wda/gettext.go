package wda

import (
	"encoding/json"
	"errors"
)

type GetTextResponse struct {
	Value  string `json:"value"`
	Status int    `json:"status"`
}

func (c *Client) GetText(elementId string) (*GetTextResponse, error) {
	if (elementId == "") {
		return nil, errors.New("Blank element id")
	}
	session, err := c.getSession()
	if err != nil {
		return nil, err
	}
	res, err := c.Client.Get("/session/" + session + "/element/" + elementId + "/text")
	if err != nil {
		return nil, err
	}
	var textResp GetTextResponse
	err = json.Unmarshal(res, &textResp)
	if err != nil {
		return nil, err
	}
	return &textResp, nil
}
