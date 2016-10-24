package handlers

import (
	"errors"
	"github.com/qa-dev/go-core/response"
	"github.com/qa-dev/wda-inspector/wda"
	"log"
	"net/http"
	"github.com/qa-dev/wda-inspector/net"
)

type GetTextHandler struct {
	WdaClient *wda.Client
}

type GetTextResponse struct {
	Text string
}

func NewGetTextHandler(c *wda.Client) *GetTextHandler {
	return &GetTextHandler{WdaClient: c}
}

func (h *GetTextHandler) getTextResponse(elementId string) (*wda.GetTextResponse, error) {
	res, err := h.WdaClient.GetText(elementId)
	if err != nil {
		return nil, err
	}
	if res.Status != net.StatusOK {
		return nil, errors.New("Bad status from inspector")
	}
	return res, err
}

func (h *GetTextHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	elementId := req.FormValue("elementId")
	if elementId == "" {
		response.Json(resp, NewJsonError("Blank element id"), http.StatusInternalServerError)
		return
	}
	s, err := h.getTextResponse(elementId)
	if err != nil {
		log.Printf(err.Error())
		response.Json(resp, NewJsonError(err.Error()), http.StatusInternalServerError)
		return
	}
	response.Json(resp, s, http.StatusOK)
}
