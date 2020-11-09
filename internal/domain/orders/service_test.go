package orders

import (
	"apitest/internal/domain/model"
	"apitest/internal/tools/errors"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initTest(err error) (context.Context, Service) {
	ctx := context.Background()
	service := NewService(fakeOrderRepository{err: err})

	return ctx, service
}

type fakeOrderRepository struct {
	err error
}

func (f fakeOrderRepository) GetById(ctx context.Context, id string) (*model.Order, error) {
	if f.err != nil {
		return nil, f.err
	}
	return &model.Order{}, nil
}

func TestService_ValidateOrder(t *testing.T) {
	ctx, service := initTest(nil)

	shipment := &model.Shipment{OrderId: "123456789"}
	err := service.ValidateOrder(ctx, shipment)
	assert.Equal(t, "", shipment.Id)
	assert.Nil(t, err)
}

func TestService_ValidateOrderError(t *testing.T) {
	ctx, service := initTest(errors.NewBusinessError("error_validating_order"))

	shipment := &model.Shipment{OrderId: "123456789"}
	err := service.ValidateOrder(ctx, shipment)
	assert.NotNil(t, err)
	assert.Equal(t, "error_validating_order", err.Error())
}

func TestService_ValidateOrderNotFound(t *testing.T) {
	ctx, service := initTest(errors.NewRestError("not_found", 404))

	shipment := &model.Shipment{OrderId: "123456789"}
	err := service.ValidateOrder(ctx, shipment)
	assert.NotNil(t, err)
	assert.Equal(t, "order does not exist", err.Error())
}
