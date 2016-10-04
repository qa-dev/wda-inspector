package handlers

import (
	"net/http"
)

type StaticHandler struct {
}

func NewStaticHandler() *StaticHandler {
	return &StaticHandler{}
}

func (h *StaticHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	path := "./resources" + req.URL.Path
	http.ServeFile(resp, req, path)
}
