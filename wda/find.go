package wda

import (
	"encoding/json"
)

type FindResponse struct {
	Value struct {
		ElementId string `json:"ELEMENT"`
		Type      string `json:"type"`
	} `json:"value"`
	Status int
}

type FindRequest struct {
	Using string `json:"using"`
	Value string `json:"value"`
}

func (c *Client) Find(using string, value string) (*FindResponse, error) {
	session, err := c.getSession()
	if err != nil {
		return nil, err
	}
	findReq := FindRequest{Using: using, Value: value}
	reqBody, err := json.Marshal(findReq)
	if err != nil {
		return nil, err
	}
	res, err := c.Client.post("/session/"+session+"/element", reqBody)
	if err != nil {
		return nil, err
	}
	var findResp *FindResponse
	err = json.Unmarshal(res, &findResp)
	if err != nil {
		return nil, err
	}
	return findResp, nil
}
