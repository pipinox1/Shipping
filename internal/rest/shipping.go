package rest

import (
	"apitest/internal/domain/model"
	"apitest/internal/tools/errors"
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

type Service interface {
	CreateShipping(ctx context.Context, shipping *model.Shipment) error
	GetShipping(ctx context.Context, id string) (*model.Shipment, error)
}

type ShippingHandler struct {
	service Service
}

func NewShippingHandler(service Service) *ShippingHandler {
	return &ShippingHandler{
		service: service,
	}
}

func (sh *ShippingHandler) createShipping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	shipping := &model.Shipment{}
	err := json.NewDecoder(r.Body).Decode(shipping)
	if err != nil {
		return
	}
	err = sh.service.CreateShipping(ctx, shipping)
	if err != nil {
		ErrorResponse(w, 400, err.Error())
		return
	}
	WebResponse(w, 201, shipping)
}

func (sh *ShippingHandler) getShipping(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		ErrorResponse(w, 400, "id is mandatory")
		return
	}
	ctx := r.Context()
	shipment, err := sh.service.GetShipping(ctx, id)
	if err != nil {
		if err == errors.NotFound {
			ErrorResponse(w, 404, "not_found")
		}
		ErrorResponse(w, 503, err.Error())
	}
	WebResponse(w, 200, shipment)
}
