package wda

import (
	"encoding/json"
)

type ElementSource struct {
	IsEnabled string           `json:"isEnabled"`
	IsVisible string           `json:"isVisible"`
	Frame     string           `json:"frame"`
	Children  []*ElementSource `json:"children"`
	Rect      struct {
		Origin struct {
			X float32 `json:"x"`
			Y float32 `json:"y"`
		} `json:"origin"`
		Size struct {
			Width  float32 `json:"width"`
			Height float32 `json:"height"`
		} `json:"size"`
	} `json:"rect"`
	Value         interface{} `json:"value"`
	Label         string      `json:"label"`
	Type          string      `json:"type"`
	Name          string      `json:"name"`
	RawIdentifier string      `json:"rawIdentifier"`
}

type Source struct {
	Value struct {
		Tree *ElementSource `json:"tree"`
	} `json:"value"`
	Status int `json:"status"`
}

func (c *Client) Source() (*Source, error) {
	res, err := c.get("/source")
	if err != nil {
		return nil, err
	}
	var s *Source
	err = json.Unmarshal(res, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
