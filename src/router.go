package main

import (
	// import a kafka struct defined by it's interface
	// import a postgres struct defined by an interface

	"github.com/rochinworks/go-home/pg"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func httpRouter(pg pg.Controller) (chi.Router, error) {
	// chi router is easy to use and lightweight
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// base routes
	r.Get("/", baseHandler())

	return r, nil
}
