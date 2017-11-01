package wda

import (
	"encoding/json"
	"errors"
)

type GetTypeResponse struct {
	Value  string `json:"value"`
	Status int    `json:"status"`
}

func (c *Client) GetType(elementId string) (*GetTypeResponse, error) {
	if elementId == "" {
		return nil, errors.New("Blank element id")
	}
	session, err := c.getSession()
	if err != nil {
		return nil, err
	}
	res, err := c.get("/session/" + session + "/element/" + elementId + "/attribute/type")
	if err != nil {
		return nil, err
	}
	var typeResp GetTypeResponse
	err = json.Unmarshal(res, &typeResp)
	if err != nil {
		return nil, err
	}
	return &typeResp, nil
}
