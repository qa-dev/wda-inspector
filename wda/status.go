package wda

import "encoding/json"

type Status struct {
	SessionId string `json:"sessionId,omitempty"`
	Status    int    `json:"status"`
}

func (c *Client) status() (*Status, error) {
	res, err := c.Client.Get("/status")
	if err != nil {
		return nil, err
	}
	var s *Status
	err = json.Unmarshal(res, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
