package shipping

import (
	"apitest/internal/domain/model"
	"apitest/internal/tools/errors"
	"context"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func initTest(err error) (context.Context, Service) {
	ctx := context.Background()
	service := NewService(&fakeShippingRepository{err: err}, &fakeOrderService{err}, fakeUserService{err})

	return ctx, service
}

type fakeShippingRepository struct {
	err error
}

func (f fakeShippingRepository) GetById(ctx context.Context, id string) (*model.Shipment, error) {
	if f.err != nil {
		return nil, f.err
	}
	return &model.Shipment{}, nil
}

func (f *fakeShippingRepository) Save(ctx context.Context, shipment *model.Shipment) error {
	return f.err
}

type fakeOrderService struct {
	err error
}

func (f *fakeOrderService) ValidateOrder(ctx context.Context, shipment *model.Shipment) error {
	if f.err != nil {
		if strings.Contains(f.err.Error(),"order"){
			return f.err
		}
	}
	return nil
}

type fakeUserService struct {
	err error
}

func (f fakeUserService) ValidateUser(ctx context.Context, shipment *model.Shipment) error {
	if f.err != nil {
		if strings.Contains(f.err.Error(),"user"){
			return f.err
		}
	}
	return f.err
}

func TestService_CreateShippingWithInvalidCost(t *testing.T) {
	ctx, service := initTest(nil)
	err := service.CreateShipping(ctx,&model.Shipment{Cost: 10})
	assert.NotNil(t,err)
	assert.Equal(t,"invalid cost",err.Error())
}


func TestService_CreateShippingWithInvalidDistributor(t *testing.T) {
	ctx, service := initTest(nil)
	err := service.CreateShipping(ctx,&model.Shipment{Cost: 1000,Distributor: "test"})
	assert.NotNil(t,err)
	assert.Equal(t,"invalid distributor",err.Error())
}

func TestService_CreateShippingWithUserError(t *testing.T) {
	ctx, service := initTest(errors.NewBusinessError("error_user"))
	err := service.CreateShipping(ctx,&model.Shipment{Cost: 1000,Distributor: "andreani"})
	assert.NotNil(t,err)
	assert.Equal(t,"error_user",err.Error())
}


func TestService_CreateShippingWithOrderError(t *testing.T) {
	ctx, service := initTest(errors.NewBusinessError("error_order"))
	err := service.CreateShipping(ctx,&model.Shipment{Cost: 1000,Distributor: "andreani"})
	assert.NotNil(t,err)
	assert.Equal(t,"error_order",err.Error())
}


