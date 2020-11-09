package user

import (
	"apitest/internal/domain/model"
	"apitest/internal/tools/restclient"
	"context"
)

type Repository interface {
	GetById(ctx context.Context, id int64) (*model.User, error)
}

func NewRepository(restClient restclient.RestClient) Repository {
	return &repository{}
}

type repository struct {
	restclient.RestClient
}

func (r *repository) GetById(ctx context.Context, id int64) (*model.User, error) {
	return &model.User{Id: "123", Name: "Eduardo", LastName: "Peña", AllowShipment: true}, nil
}
