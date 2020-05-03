package handlers

import "net/http"

type Home struct {
}

func NewHome() *Home {
	return &Home{}
}

func (h *Home) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`shortner`))
}
