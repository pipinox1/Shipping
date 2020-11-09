package orders

import (
	"apitest/internal/domain/model"
	"apitest/internal/tools/restclient"
	"context"
	"fmt"
)

var requestUlr = "http://localhost:3004/v1/orders/%s"

type Repository interface {
	GetById(ctx context.Context, id string) (*model.Order, error)
}

func NewRepository(restClient restclient.RestClient) Repository {
	return &repository{
		restClient,
	}
}

type repository struct {
	restclient.RestClient
}

func (r *repository) GetById(ctx context.Context, id string) (*model.Order, error) {

	url := fmt.Sprintf(requestUlr, id)
	order := &model.Order{}
	err := r.DoGet(ctx, url, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}
