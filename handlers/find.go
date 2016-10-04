package handlers

import (
	"errors"
	"github.com/qa-dev/wda-inspector/element"
	"github.com/qa-dev/wda-inspector/wda"
	"log"
	"net/http"
	"github.com/qa-dev/go-core/response"
)

type FindHandler struct {
	WdaClient *wda.Client
}

type FindResponse struct {
	*wda.RectResponse
}

func NewFindHandler(c *wda.Client) *FindHandler {
	return &FindHandler{WdaClient: c}
}

func (h *FindHandler) find(using string, value string) (*wda.FindResponse, error) {
	res, err := h.WdaClient.Find(using, value)
	if err != nil {
		return nil, err
	}
	if res.Status != wda.StatusOK {
		return nil, errors.New("Bad status from inspector")
	}
	return res, err
}

func (h *FindHandler) rect(elId string) (*wda.RectResponse, error) {
	res, err := h.WdaClient.Rect(elId)
	if err != nil {
		return nil, err
	}
	if res.Status != wda.StatusOK {
		return nil, errors.New("Bad status from inspector")
	}
	return res, err
}

func (h *FindHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	f, err := h.find(req.FormValue("using"), req.FormValue("value"))
	if err != nil {
		log.Printf(err.Error())
		response.Json(resp, NewJsonError(err.Error()), http.StatusInternalServerError)
		return
	}
	if f.Value.Type == element.TypeOther {
		response.Json(resp, NewJsonError("Element not found on page"), http.StatusBadRequest)
		return
	}

	r, err := h.rect(f.Value.ElementId)
	if err != nil {
		log.Printf(err.Error())
		response.Json(resp, NewJsonError(err.Error()), http.StatusInternalServerError)
		return
	}
	response.Json(resp, r, http.StatusOK)
}
