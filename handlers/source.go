package handlers

import (
	"errors"
	"github.com/qa-dev/go-core/response"
	"github.com/qa-dev/wda-inspector/wda"
	"log"
	"net/http"
)

type SourceHandler struct {
	WdaClient *wda.Client
}

type SourceResponse struct {
	Tree *wda.ElementSource `json:"tree"`
}

func NewSourceHandler(c *wda.Client) *SourceHandler {
	return &SourceHandler{WdaClient: c}
}

func (h *SourceHandler) source() (*wda.Source, error) {
	sc, err := h.WdaClient.Source()
	if err != nil {
		return nil, err
	}
	if sc.Status != wda.StatusOK {
		return nil, errors.New("Bad status from inspector")
	}
	return sc, err
}

func (h *SourceHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	s, err := h.source()
	if err == nil {
		data := &SourceResponse{Tree: s.Value.Tree}
		response.Json(resp, data, http.StatusOK)
	} else {
		log.Printf(err.Error())
		response.Json(resp, NewJsonError(err.Error()), http.StatusInternalServerError)
	}
}
