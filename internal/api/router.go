package api

import (
	"compress/flate"
	"github.com/douglasfsti/golang-shortener-api/config"
	"github.com/douglasfsti/golang-shortener-api/internal/api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	Home     = "/"
	Shortner = "/v1/shortner"
	Redirect = "/v1/{code}"
)

func GetRouter(router *chi.Mux, container config.Container) *chi.Mux {
	EnableMiddleware(router)
	ConfigureRoutes(router, container)

	return router
}

func ConfigureRoutes(router chi.Router, container config.Container) {
	router.Route(Home, func(r chi.Router) {
		handler := handlers.NewHome()
		r.Get("/", handler.Get)
	})

	router.Route(Shortner, func(r chi.Router) {
		handler := handlers.NewShortner(container.GetShortnerService())
		r.Post("/", handler.Post)
	})

	router.Route(Redirect, func(r chi.Router) {
		handler := handlers.NewRedirect(container.GetShortnerService())
		r.Get("/", handler.Get)
	})
}

func EnableMiddleware(router chi.Router) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(flate.BestCompression))
	router.Use(middleware.Timeout(config.HTTPServerTimeout))
}
