package wda

import (
	"encoding/json"
	"fmt"
)

type FindSuccessResponse struct {
	Value struct {
		ElementId string `json:"ELEMENT"`
		Type      string `json:"type"`
	} `json:"value"`
	Status int
}

type FindErrorResponse struct {
	Value struct {
		Using      string `json:"using"`
		Value string `json:"value"`
		Description string `json:"description"`
	} `json:"value"`
	Status int
}

type FindRequest struct {
	Using string `json:"using"`
	Value string `json:"value"`
}

func (c *Client) Find(using string, value string) (*FindSuccessResponse, error) {
	session, err := c.getSession()
	if err != nil {
		return nil, err
	}
	findReq := FindRequest{Using: using, Value: value}
	reqBody, err := json.Marshal(findReq)
	if err != nil {
		return nil, err
	}
	res, err := c.post("/session/"+session+"/element", reqBody)
	if err != nil {
		return nil, err
	}
	var findResp *FindSuccessResponse
	err = json.Unmarshal(res, &findResp)
	if err != nil {
		return nil, err
	}
	if findResp.Status != StatusOK {
		var findErrorResp *FindErrorResponse
		findErrorRespErr := json.Unmarshal(res, &findErrorResp)
		if findErrorRespErr == nil {
			return nil, fmt.Errorf("WDA returns error on finding element: %+v", findErrorResp)
		}
	}
	return findResp, nil
}
