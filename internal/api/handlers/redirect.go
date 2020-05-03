package handlers

import (
	"github.com/douglasfsti/golang-shortener-api/pkg/shortner"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type Redirect struct {
	Service shortner.Service
}

func NewRedirect(service shortner.Service) *Redirect {
	return &Redirect{
		Service: service,
	}
}

func (rd *Redirect) Get(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`missing code`))
		return
	}

	redirect, err := rd.Service.Find(code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	log.Printf("redirecting %v to %v", code, redirect.URL)
	http.Redirect(w, r, redirect.URL, http.StatusMovedPermanently)
}
