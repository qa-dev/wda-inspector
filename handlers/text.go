package handlers

import (
	"errors"
	"github.com/qa-dev/go-core/response"
	"github.com/qa-dev/wda-inspector/wda"
	"log"
	"net/http"
)

type TextHandler struct {
	WdaClient *wda.Client
}

type TextResponse struct {
	Text string
}

func NewTextHandler(c *wda.Client) *TextHandler {
	return &TextHandler{WdaClient: c}
}

func (h *TextHandler) get(elementId string) (*wda.TextResponse, error) {
	res, err := h.WdaClient.GetText(elementId)
	if err != nil {
		return nil, err
	}
	if res.Status != wda.StatusOK {
		return nil, errors.New("Bad status from inspector")
	}
	return res, err
}

func (h *TextHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	s, err := h.get(req.FormValue("elementId"))
	if err != nil {
		log.Printf(err.Error())
		response.Json(resp, NewJsonError(err.Error()), http.StatusInternalServerError)
		return
	}
	//if s.Value.Type == element.TypeOther {
	//	response.Json(resp, NewJsonError("Element not found on page"), http.StatusBadRequest)
	//	return
	//}
	// TODO
}
