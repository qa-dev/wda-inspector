package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type IndexPageHandler struct {
}

func NewIndexPage() *IndexPageHandler {
	return &IndexPageHandler{}
}

func (h *IndexPageHandler) index() (*template.Template, error) {
	t := template.New("index.html")
	return t.ParseFiles("resources/templates/index.html")
}

func (h *IndexPageHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	t, err := h.index()
	if err != nil {
		log.Printf(err.Error())
	}
	var data interface{}
	err = t.Execute(resp, data)
	if err != nil {
		log.Printf(err.Error())
	}
}
