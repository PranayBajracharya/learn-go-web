package utils

import "net/http"

type JsonResponse struct {
	http.ResponseWriter
}

func (w *JsonResponse) Write(b []byte) (int, error) {
	w.Header().Set("Content-Type", "application/json")
	return w.ResponseWriter.Write(b)
}

func NewJsonResponse(w http.ResponseWriter) *JsonResponse {
	return &JsonResponse{w}
}
