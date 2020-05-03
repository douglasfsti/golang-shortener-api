package handlers

import (
	"github.com/douglasfsti/golang-shortener-api/pkg/shortner"
	"net/http"
)

type Shortner struct {
	Service shortner.Service
}

func NewShortner(service shortner.Service) *Shortner {
	return &Shortner{
		Service: service,
	}
}

func (s *Shortner) Post(w http.ResponseWriter, r *http.Request) {
	redirect, err := s.Service.NewRedirectIoReader(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err = s.Service.Store(redirect)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	response, err := s.Service.Encode(redirect)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}
