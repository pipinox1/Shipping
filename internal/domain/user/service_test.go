package user

import (
	"apitest/internal/domain/model"
	"apitest/internal/tools/errors"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type userRepositoryMock struct {
	mock.Mock
}

func (r *userRepositoryMock) GetById(ctx context.Context, userId int64) (*model.User, error) {
	args := r.Called(mock.Anything)
	user, _ := args.Get(0).(*model.User)
	return user, args.Error(1)
}

func initTest() (context.Context, *userRepositoryMock, Service) {
	userRepo := &userRepositoryMock{}
	ctx := context.Background()
	service := NewService(userRepo)
	return ctx, userRepo, service
}

func TestService_ValidateUserNotFound(t *testing.T) {
	ctx, userRepo, service := initTest()
	userRepo.Mock.On("GetById").Return(nil, errors.NewRestError("not_found", 404))
	err := service.ValidateUser(ctx, &model.Shipment{})
	assert.NotNil(t, err)
	assert.Equal(t, "not_found", err.Error())
}

func TestService_ValidateUserUserNotAllowedToShipProduct(t *testing.T) {
	ctx, userRepo, service := initTest()
	userRepo.Mock.On("GetById").Return(model.User{LastName: "Perez", Name: "Juan", AllowShipment: false}, nil)
	err := service.ValidateUser(ctx, &model.Shipment{})
	assert.NotNil(t, err)
	assert.Equal(t, "The user does not have permission to send products", err.Error())
}

func TestService_ValidateUserOk(t *testing.T) {
	ctx, userRepo, service := initTest()
	userRepo.Mock.On("GetById").Return(model.User{LastName: "Perez", Name: "Juan", AllowShipment: true}, nil)
	err := service.ValidateUser(ctx, &model.Shipment{})
	assert.Nil(t, err)
}
