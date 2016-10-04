package handlers

import (
	"errors"
	"github.com/qa-dev/wda-inspector/wda"
	"log"
	"net/http"
	"github.com/qa-dev/go-core/response"
)

type ScreenshotHandler struct {
	WdaClient *wda.Client
}

type ScreenshotResponse struct {
	Img string `json:"img"`
}

func NewScreenshotHandler(c *wda.Client) *ScreenshotHandler {
	return &ScreenshotHandler{WdaClient: c}
}

func (h *ScreenshotHandler) screenshot() (*wda.Screenshot, error) {
	sc, err := h.WdaClient.Screenshot()
	if err != nil {
		return nil, err
	}
	if sc.Status != wda.StatusOK {
		return nil, errors.New("Bad status from inspector")
	}
	return sc, err
}

func (h *ScreenshotHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	sc, err := h.screenshot()
	if err == nil {
		data := &ScreenshotResponse{Img: sc.Value}
		response.Json(resp, data, http.StatusOK)
	} else {
		log.Printf(err.Error())
		response.Json(resp, NewJsonError(err.Error()), http.StatusInternalServerError)
	}
}
