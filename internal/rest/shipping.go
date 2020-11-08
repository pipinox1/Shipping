package rest

import (
	"apitest/internal/domain/model"
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	CreateShipping(ctx context.Context,shipping *model.Shipment) error
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
	err = sh.service.CreateShipping(ctx,shipping)
	if err != nil {
		ErrorResponse(w, 400, err.Error())
	}
	WebResponse(w, 201, shipping)
}

func (sh *ShippingHandler) getShipping(w http.ResponseWriter, r *http.Request) {
	WebResponse(w,200,"respuesta al get")
}
