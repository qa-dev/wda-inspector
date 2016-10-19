package wda

import (
	"encoding/json"
	"log"
)

type SessionResponse struct {
	Value struct {
		SessionId string `json:"sessionId"`
	} `json:"value"`
	Status int `json:"status"`
}

func (c *Client) session() (*SessionResponse, error) {
	req := make(map[string]map[string]string, 1)
	tmpMap := make(map[string]string, 1)
	tmpMap["bundleId"] = c.BundleId()
	req["desiredCapabilities"] = tmpMap
	reqRaw, err := json.Marshal(req)
	log.Printf(string(reqRaw))
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Post("/session", reqRaw)
	log.Printf(string(res))
	if err != nil {
		return nil, err
	}
	var session *SessionResponse
	err = json.Unmarshal(res, &session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (c *Client) getSession() (string, error) {
	status, err := c.status()
	if err != nil {
		return "", err
	}
	if status.SessionId == "" {
		res, err := c.session()
		if err != nil {
			return "", err
		}
		return res.Value.SessionId, nil
	} else {
		return status.SessionId, nil
	}
}
