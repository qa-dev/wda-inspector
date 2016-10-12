package wda

import (
	"encoding/json"
)

type GetTextResponse struct {
	Value  string `json:"value"`
	Status int    `json:"status"`
}

func (c *Client) GetText(elementId string) (*GetTextResponse, error) {
	session, err := c.getSession()
	if err != nil {
		return nil, err
	}
	res, err := c.get("/session/" + session + "/element/" + elementId + "/text")
	if err != nil {
		return nil, err
	}
	var textResp *GetTextResponse
	err = json.Unmarshal(res, &textResp)
	if err != nil {
		return nil, err
	}
	return textResp, nil
}
