package wda

import (
	"encoding/json"
)

type TextResponse struct {
	Value struct {
		Text string `json:"text"`
	} `json:"value"`
	Status int `json:"status"`
}

type TextRequest struct {
	ElementId string `json:"elementId"`
}

func (c *Client) GetText(elementId string) (*TextResponse, error) {
	session, err := c.getSession()
	if err != nil {
		return nil, err
	}
	textReq := TextRequest{ElementId:elementId}
	reqBody, err := json.Marshal(textReq)
	if err != nil {
		return nil, err
	}
	res, err := c.post("/session/"+session+"/element/"+elementId+"/text", reqBody)
	if err != nil {
		return nil, err
	}
	var textResp *TextResponse
	err = json.Unmarshal(res, &textResp)
	if err != nil {
		return nil, err
	}
	return textResp, nil
}
