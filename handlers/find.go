package handlers

import (
	"errors"
	"github.com/qa-dev/go-core/response"
	"github.com/qa-dev/wda-inspector/element"
	"github.com/qa-dev/wda-inspector/wda"
	"log"
	"net/http"
	"strings"
)

type FindHandler struct {
	WdaClient *wda.Client
}

type FindResponse struct {
	*wda.RectResponse
}

type TypeResponse struct {
	*wda.GetTypeResponse
}

type Response struct {
	Value struct {
		X      float32 `json:"x"`
		Y      float32 `json:"y"`
		Width  float32 `json:"width"`
		Height float32 `json:"height"`
	} `json:"value"`
	Type   string  `json:"type"`
	Status int `json:"status"`
}

func NewFindHandler(c *wda.Client) *FindHandler {
	return &FindHandler{WdaClient: c}
}

func (h *FindHandler) find(using string, value string) (*wda.FindSuccessResponse, error) {
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

func (h *FindHandler) typ(elId string) (*wda.GetTypeResponse, error) {
	typ, err := h.WdaClient.GetType(elId)
	if err != nil {
		return nil, err
	}
	if typ.Status != wda.StatusOK {
		return nil, errors.New("Bad status from inspector")
	}
	return typ, err
}

func (h *FindHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	f, err := h.find(req.FormValue("using"), req.FormValue("value"))
	if err != nil {
		log.Printf(err.Error())
		response.Json(resp, NewJsonError(err.Error()), http.StatusInternalServerError)
		return
	}

	r, err := h.rect(f.Value.ElementId)
	if err != nil {
		log.Printf(err.Error())
		response.Json(resp, NewJsonError(err.Error()), http.StatusInternalServerError)
		return
	}

	t, err := h.typ(f.Value.ElementId)
	if err != nil {
		log.Printf(err.Error())
		response.Json(resp, NewJsonError(err.Error()), http.StatusInternalServerError)
		return
	}

	var res Response
	res.Type = strings.Replace(t.Value, "XCUIElementType", "", -1)
	res.Value.Height = r.Value.Height
	res.Value.Width = r.Value.Width
	res.Value.X = r.Value.X
	res.Value.Y = r.Value.Y
	res.Status = res.Status

	if res.Type == element.TypeOther && r.IsInvalid() {
		response.Json(resp, NewJsonError("Element not found on page"), http.StatusBadRequest)
		return
	}
	response.Json(resp, res, http.StatusOK)
}
