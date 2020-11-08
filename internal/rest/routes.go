package rest

import (
	"github.com/go-chi/chi"
)

func RegisterShippingRoute(r chi.Router, handler *ShippingHandler) {
	r.Use(securityMiddleware)
	r.Route("/shipping", func(r chi.Router) {
		r.Get("/{id}", handler.getShipping)
		r.Post("/", handler.createShipping)
	})
}
