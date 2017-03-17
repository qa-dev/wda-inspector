package wda

import (
	"encoding/json"
)

type RectResponse struct {
	Value struct {
		X      float32 `json:"x"`
		Y      float32 `json:"y"`
		Width  float32 `json:"width"`
		Height float32 `json:"height"`
	} `json:"value"`
	Status int `json:"status"`
}

func (c *Client) Rect(elId string) (*RectResponse, error) {
	session, err := c.getSession()
	if err != nil {
		return nil, err
	}
	res, err := c.get("/session/" + session + "/element/" + elId + "/rect")
	if err != nil {
		return nil, err
	}
	var rectResp *RectResponse
	err = json.Unmarshal(res, &rectResp)
	if err != nil {
		return nil, err
	}
	return rectResp, nil
}

// IsInvalid проверяет, валидный ли элемент, если вдруг вместо нужного и
// несуществующего элемента вернулся XCUIElementTypeElementOther, проверяя значения его /rect.
// Если все нули – значит элемент не валиден
func (r *RectResponse) IsInvalid() bool {
	return r.Value.Height == 0 && r.Value.Width == 0 && r.Value.X == 0 && r.Value.Y == 0
}
