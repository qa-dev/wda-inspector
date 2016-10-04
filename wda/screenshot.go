package wda

import (
	"encoding/json"
)

type Screenshot struct {
	Value  string `json:"value"`
	Status int    `json:"status"`
}

func (c *Client) Screenshot() (*Screenshot, error) {
	res, err := c.get("/screenshot")
	if err != nil {
		return nil, err
	}
	var ss *Screenshot
	err = json.Unmarshal(res, &ss)
	if err != nil {
		return nil, err
	}
	return ss, nil
}
