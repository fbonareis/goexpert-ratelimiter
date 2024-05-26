package web

import "net/http"

type WebSampleHandler struct {
}

func NewWebSampleHandler() *WebSampleHandler {
	return &WebSampleHandler{}
}

func (h *WebSampleHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sample endpoint response!"))
}
