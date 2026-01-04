package handlers

import (
	"github.com/Josesx506/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		// Local middleware for account route
		router.Use(middleware.Authorization)

		// Embedded route - /account/coins endpoint
		router.Get("/coins", GetCoinBalance)
	})
}
